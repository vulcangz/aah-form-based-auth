The code was written for [this article](https://www.worldlink.com.cn/post/understand-the-security-design-of-go-aah-framework-and-application-examples.html).

<p align="center">
  <img src="https://cdn.aahframework.org/assets/img/aah-logo-64x64.png" />
  <h2 align="center">Example - Form based Auth using a database</h2>
</p>

This example is based on the official example ([Example - Form based Auth](https://github.com/go-aah/examples/tree/v0.12.x/form-based-auth)) but uses a MySQL database as backend storage. Form based auth includes authentication, route authorization via routes config, access permissions on view files and session access on View files.

Learn more about [Security design](https://docs.aahframework.org/security-design.html), [Authentication](https://docs.aahframework.org/authentication.html), [Authorization](https://docs.aahframework.org/authorization.html) and [Form Auth scheme](https://docs.aahframework.org/auth-schemes/form.html).

## Usage


### Database installation

Create MySQL database and import data

The db schma with demo data is located at [docs/dbschema/schema.sql](docs/dbschema/schema.sql)

### Database configuration

edit config file: [config/aah.conf](config/aah.conf)

```bash
database {
  driver = "mysql"
  host = "localhost"
  port = "3306"
  username = "root"  # change to yours
  password = "mysql" # change to yours
  name = "aah-form-based-auth"  # change to your db
  max_idle_connections = 20
  max_active_connections = 30
  max_connection_lifetime = 2
}
```

```bash
git clone https://github.com/vulcangz/aah-form-based-auth.git
```

### Run this example

```bash
cd aah-form-based-auth
aah run
```

### Visit this URL

  * http://localhost:8080

The application will take you to the login page. From there, it is self explanatory. Happy coding!

Navigate these URLs with various credentials listed on the login page. Observe the application logs to learn more.

  * http://localhost:8080/manage/users.html
  * http://localhost:8080/admin/dashboard.html
  * http://localhost:8080/login.html


## Credits and Inspiration

* [Svelte framework](https://svelte.dev/)
* [SQLBoiler](https://github.com/volatiletech/sqlboiler)