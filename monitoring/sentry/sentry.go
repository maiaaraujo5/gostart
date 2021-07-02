package sentry

import (
	"github.com/getsentry/sentry-go"
	"github.com/maiaaraujo5/gostart/log/logger"
	"github.com/maiaaraujo5/gostart/monitoring"
)

func Init() error {
	config, err := monitoring.NewConfig()
	if err != nil {
		return err
	}

	err = sentry.Init(sentry.ClientOptions{
		Dsn:              config.AuthenticationKey,
		AttachStacktrace: true,
		TracesSampleRate: 1,
		Release:          config.Release,
		Environment:      config.Environment,
	})

	if err != nil {
		return err
	}
	logger.Info("sentry successfully configured")
	return nil
}
