package sentry

import (
	"github.com/getsentry/sentry-go"
	"github.com/maiaaraujo5/gostart/monitoring"
)

type Sentry struct {
	Sentry *sentry.Client
}

func NewSentry(config *monitoring.Config) (*Sentry, error) {
	client, err := sentry.NewClient(sentry.ClientOptions{
		Dsn:         config.AuthenticationKey,
		Release:     config.Release,
		Environment: config.Environment,
	})

	if err != nil {
		return nil, err
	}

	return &Sentry{
		Sentry: client,
	}, nil
}
