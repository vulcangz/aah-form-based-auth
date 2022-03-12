package queue

import (
	"encoding/json"
	"fmt"

	aah "aahframe.work"
	"github.com/streadway/amqp"
)

// Code borrowed from the examples of "github.com/streadway/amqp"
// @see https://github.com/streadway/amqp/tree/master/_examples

var (
	AMQPURI      = "amqp://guest:guest@localhost:5672/" // AMQP URI
	ExchangeName = "aah.users"                          // Durable, non-auto-deleted AMQP exchange name
	ExchangeType = "direct"                             // Exchange type - direct|fanout|topic|x-custom
	RoutingKey   = "users-route"                        // AMQP binding key
	QueueName    = "user.update"                        // Ephemeral AMQP queue name
	ConsumerTag  = "ppv-consumer"                       // AMQP consumer tag (should not be blank)
)

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	tag     string
	done    chan error
}

func NewConsumer(ctx *aah.Context, m *aah.Middleware) (*Consumer, error) {
	c := &Consumer{
		conn:    nil,
		channel: nil,
		tag:     ConsumerTag,
		done:    make(chan error),
	}

	var err error

	aah.App().Log().Debugf("dialing %q", AMQPURI)
	c.conn, err = amqp.Dial(AMQPURI)
	if err != nil {
		return nil, fmt.Errorf("dial: %s", err)
	}

	go func() {
		aah.App().Log().Debugf("closing: %s", <-c.conn.NotifyClose(make(chan *amqp.Error)))
	}()

	aah.App().Log().Debugf("got Connection, getting Channel")
	c.channel, err = c.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("channel: %s", err)
	}

	aah.App().Log().Debugf("got Channel, declaring Exchange (%q)", ExchangeName)
	if err = c.channel.ExchangeDeclare(
		ExchangeName, // name of the exchange
		ExchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return nil, fmt.Errorf("ExchangeName Declare: %s", err)
	}

	aah.App().Log().Debugf("declared Exchange, declaring Queue %q", QueueName)
	queue, err := c.channel.QueueDeclare(
		QueueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("queue Declare: %s", err)
	}

	aah.App().Log().Debugf("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
		queue.Name, queue.Messages, queue.Consumers, RoutingKey)

	if err = c.channel.QueueBind(
		queue.Name,   // name of the queue
		RoutingKey,   // bindingKey
		ExchangeName, // sourceExchange
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return nil, fmt.Errorf("queue Bind: %s", err)
	}

	aah.App().Log().Debugf("Queue bound to Exchange, starting Consume (consumer tag %q)", c.tag)
	deliveries, err := c.channel.Consume(
		queue.Name, // name
		c.tag,      // consumerTag,
		false,      // noAck
		false,      // exclusive
		false,      // noLocal
		false,      // noWait
		nil,        // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("queue Consume: %s", err)
	}

	go handle(ctx, m, deliveries, c.done)

	return c, nil
}

func (c *Consumer) Shutdown() error {
	// will close() the deliveries channel
	if err := c.channel.Cancel(c.tag, true); err != nil {
		return fmt.Errorf("consumer cancel failed: %s", err)
	}

	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}

	defer aah.App().Log().Debugf("AMQP shutdown OK")

	// wait for handle() to exit
	return <-c.done
}

func handle(ctx *aah.Context, m *aah.Middleware, deliveries <-chan amqp.Delivery, done chan error) {
	for d := range deliveries {
		aah.App().Log().Debugf(
			"got %dB delivery: [%v] %q",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)

		ppv := ctx.Subject().PrimaryPrincipal().Value
		aah.App().Log().Debugf("current subject primary principal value=%v", ppv)

		var ue UserEvent
		err := json.Unmarshal(d.Body, &ue)
		if err != nil {
			ctx.Log().Debugf("error decoding Json: %s", err)
		}

		u := ue.User
		ctx.Log().Debugf("current ppv=%v updated user's email=%v", ppv, u.Email)
		if u.Email == ppv {
			// it's me, need to be updated
			ctx.Log().Debug("oh, this needs to be updated...")
			ctx.Session().Set("FirstName", u.FirstName)
			ctx.Session().Set("LastName", u.LastName)
			ctx.Session().Set("Email", u.Email)
			ctx.Session().Set("Roles", ue.Roles)
			ctx.Session().Set("Perms", ue.Perms)

			ctx.AddViewArg("Roles", ue.Roles)
			ctx.AddViewArg("Perms", ue.Perms)
			d.Ack(false)
		} else {
			// not me, need to be requeue
			d.Nack(false, true)
		}
	}
	aah.App().Log().Debugf("handle: deliveries channel closed")
	done <- nil
}

func Publish(body []byte, reliable bool) error {

	// This function dials, connects, declares, publishes, and tears down,
	// all in one go. In a real service, you probably want to maintain a
	// long-lived connection as state, and publish against that.

	aah.App().Log().Debugf("dialing %q", AMQPURI)
	connection, err := amqp.Dial(AMQPURI)
	if err != nil {
		return fmt.Errorf("dial: %s", err)
	}
	defer connection.Close()

	aah.App().Log().Debugf("got Connection, getting Channel")
	channel, err := connection.Channel()
	if err != nil {
		return fmt.Errorf("channel: %s", err)
	}

	aah.App().Log().Debugf("got Channel, declaring %q Exchange (%q)", ExchangeType, ExchangeName)
	if err := channel.ExchangeDeclare(
		ExchangeName, // name
		ExchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return fmt.Errorf("ExchangeName Declare: %s", err)
	}

	// Reliable publisher confirms require confirm.select support from the
	// connection.
	if reliable {
		aah.App().Log().Debugf("enabling publishing confirms.")
		if err := channel.Confirm(false); err != nil {
			return fmt.Errorf("channel could not be put into confirm mode: %s", err)
		}

		confirms := channel.NotifyPublish(make(chan amqp.Confirmation, 1))

		defer confirmOne(confirms)
	}

	aah.App().Log().Debugf("declared Exchange, publishing %dB body (%q)", len(body), body)
	if err = channel.Publish(
		ExchangeName, // publish to an ExchangeName
		RoutingKey,   // routing to 0 or more queues
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            body,
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		return fmt.Errorf("ExchangeName Publish: %s", err)
	}

	return nil
}

// One would typically keep a channel of publishings, a sequence number, and a
// set of unacknowledged sequence numbers and loop until the publishing channel
// is closed.
func confirmOne(confirms <-chan amqp.Confirmation) {
	aah.App().Log().Debugf("waiting for confirmation of one publishing")

	if confirmed := <-confirms; confirmed.Ack {
		aah.App().Log().Debugf("confirmed delivery with delivery tag: %d", confirmed.DeliveryTag)
	} else {
		aah.App().Log().Debugf("failed delivery of delivery tag: %d", confirmed.DeliveryTag)
	}
}
