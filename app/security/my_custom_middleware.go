package security

import (
	"aah-form-based-auth/app/queue"

	aah "aahframe.work"
)

//‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾
// My Custom Middleware
//______________________________________________________________________________

// MyCustomMiddleware is aah Custom Authorization Middleware.
func MyCustomMiddleware(ctx *aah.Context, m *aah.Middleware) {
	if !ctx.Subject().IsAuthenticated() {
		m.Next(ctx)
		return
	}

	// If session is authenticated then populate subject and continue the request flow.
	if ctx.Subject().IsAuthenticated() {

		ppv := ctx.Subject().PrimaryPrincipal().Value

		if ppv == "admin@aahframework.org" {
			aah.App().Log().Debugf("current subject ppv=%v, skip the processing flow", ppv)
			m.Next(ctx)
			return
		}

		aah.App().Log().Debugf("current subject ppv=%v, enter the processing flow", ppv)
		// rmq := queue.RabbitInit()
		// queue.StartConsumer(ctx, m, rmq)
		// queue.StartConsumer(ctx, m)

		c, err := queue.NewConsumer(ctx, m)
		if err != nil {
			ctx.Log().Debugf("%s", err)
		}

		if err := c.Shutdown(); err != nil {
			ctx.Log().Debugf("error during shutdown: %s", err)
		}

		// queue.ConyConsumer(ctx, m)
		// queue.HopConsumer(ctx, m)
		// queue.StartReceiverr(ctx, m)

		// go func() {
		// 	if msg, ok := <-pubsub.GetPS().C; ok {
		// 		ctx.Log().Printf("Received %s, times.\n", msg)
		// 	} else {

		// 	}
		// }()

		m.Next(ctx)
		return
	}
}
