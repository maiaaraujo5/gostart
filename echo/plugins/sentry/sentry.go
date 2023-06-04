package sentry

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	echoLibrary "github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/gostart/echo"
	"github.com/maiaaraujo5/gostart/log/logger"
)

func NewSentry(ctx context.Context) echo.Plugin {
	logger.Trace("registering sentry echo plugin")

	c, err := NewConfig()
	if err != nil {
		logger.Fatal("error to get config from echo sentry plugin")
	}

	if !c.Enabled {
		return nil
	}

	return func(ctx context.Context, client *echoLibrary.Echo) error {
		err := sentry.Init(sentry.ClientOptions{
			EnableTracing:    true,
			Dsn:              c.Dsn,
			Debug:            c.Debug,
			TracesSampleRate: c.TraceSampleRate,
			Environment:      c.Environment,
		})

		if err != nil {
			return fmt.Errorf("error to initialize echo sentry plugin: %w", err)
		}

		client.Use(sentryecho.New(sentryecho.Options{
			Repanic:         true,
			WaitForDelivery: false,
			Timeout:         c.Timeout,
		}))

		sentry.CaptureMessage("It works!")

		return nil
	}
}
