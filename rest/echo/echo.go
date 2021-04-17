package echo

import (
	echoLibrary "github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/gostart/rest"
	"github.com/maiaaraujo5/gostart/rest/handler"
	"log"
)

type echo struct {
	Config *Config
	Echo   *echoLibrary.Echo
}

func NewEcho(config *Config) rest.Rest {
	return &echo{
		Config: config,
		Echo:   echoLibrary.New(),
	}
}

func (e *echo) AddRoute(method, route string, handler handler.Handler) {
	e.Echo.Add(method, route, handler.Echo.Handler, handler.Echo.Middleware...)
}

func (e *echo) Start() {
	log.Fatal(e.Echo.Start(":" + e.Config.Port))
}
