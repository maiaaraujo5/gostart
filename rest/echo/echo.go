package echo

import (
	echoLibrary "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/maiaaraujo5/gostart/rest"
	"github.com/maiaaraujo5/gostart/rest/handler"
	"log"
	"net/http"
)

type echo struct {
	Config *rest.Config
	Echo   *echoLibrary.Echo
}

func NewEcho(client *echoLibrary.Echo) rest.Rest {
	config, err := rest.NewConfig()
	if err != nil {
		return nil
	}

	if config.Cors {
		if len(config.AllowOrigins) == 0 {
			client.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
		} else {
			client.Use(middleware.CORSWithConfig(middleware.CORSConfig{
				AllowOrigins: config.AllowOrigins,
				AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
			}))
		}

	}

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
