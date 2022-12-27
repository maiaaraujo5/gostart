package echo

import (
	"context"
	echoLibrary "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/maiaaraujo5/gostart/log/logger"
	"net/http"
)

type Plugin func(ctx context.Context, client *echoLibrary.Echo) error

func NewEcho(ctx context.Context, plugins ...Plugin) (*echoLibrary.Echo, error) {
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}

	echo := echoLibrary.New()
	cors(echo, config)

	for _, pluginFunc := range plugins {
		err := pluginFunc(ctx, echo)
		if err != nil {
			logger.Fatal(err.Error())
		}
	}

	return echo, nil

}

func cors(echo *echoLibrary.Echo, config *Config) {
	if !config.Cors {
		return
	}

	if len(config.AllowOrigins) == 0 {
		echo.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
		return
	}

	echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: config.AllowOrigins,
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
}
