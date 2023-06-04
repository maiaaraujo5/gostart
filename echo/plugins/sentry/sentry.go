package sentry

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	echoLibrary "github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/gostart/echo"
	"github.com/maiaaraujo5/gostart/log/logger"
	"time"
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

		client.Use(sentryMiddleware)

		return nil
	}
}

func sentryMiddleware(next echoLibrary.HandlerFunc) echoLibrary.HandlerFunc {
	return func(c echoLibrary.Context) error {
		req := c.Request()
		hub := sentry.CurrentHub().Clone()
		hub.Scope().SetRequest(req)

		hub.AddBreadcrumb(&sentry.Breadcrumb{
			Message: "Request",
			Data: map[string]interface{}{
				"method":     req.Method,
				"url":        req.URL,
				"user-agent": req.UserAgent(),
				"header":     req.Header,
			},
			Level:     sentry.LevelInfo,
			Timestamp: time.Now(),
		}, nil)

		err := next(c)

		if err != nil {
			hub.CaptureException(err)
		}

		return err
	}

}
