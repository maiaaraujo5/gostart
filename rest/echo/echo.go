package echo

import (
	echoLibrary "github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/gostart/rest"
	"github.com/maiaaraujo5/gostart/rest/handler"
	"log"
)

type echo struct {
	Config *rest.Config
	Echo   *echoLibrary.Echo
}

func NewEcho(config *rest.Config, client *echoLibrary.Echo) rest.Rest {
	return &echo{
		Config: config,
		Echo:   client,
	}
}

func (e *echo) AddRoute(method, route string, handler handler.Handler) {
	e.Echo.Add(method, route, handler.Echo.Handler, handler.Echo.Middleware...)
}

func (e *echo) Start() {
	log.Fatal(e.Echo.Start(":" + e.Config.Port))
}
