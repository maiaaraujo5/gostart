# *GO START*

A helper to start clients and applications.

## How to use

### Easy Way

It's very simple to use gostart, if you use fx in your application you can start the application like this:

 ```shell
       application.Run(
		application.Start.
			WithEcho().
			WithMongo().
			WithRabbitMQ().
			WithCustomProvider(fx.AuthenticatorHandlerModule()).
			Build(),
	)
```

In this example, all these providers will be initialized for you, with default configurations.

### Start client as you want

If you don't use fx, or you don't want to initialize the clients in the same time, you can call the methods to
initialize in your code like this:

 ```shell
 import (
	"github.com/maiaaraujo5/gostart/cache/redis"
	
	func main() {
	  client, err := redis.Connect()
	  if err!=nil {... //do something with error}	
	  
	  ... //do something with redis client
	  }
	
)
```