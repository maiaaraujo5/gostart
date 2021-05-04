package application

import (
	"github.com/lann/builder"
	"github.com/maiaaraujo5/gostart/config"
	mongo "github.com/maiaaraujo5/gostart/database/mongodb/fx"
	sentry "github.com/maiaaraujo5/gostart/monitoring/sentry/fx"
	echo "github.com/maiaaraujo5/gostart/rest/echo/fx"
	"go.uber.org/fx"
)

type Options struct {
	Providers []fx.Option
}

type providersBuilder builder.Builder

func (b providersBuilder) WithEcho() providersBuilder {
	return builder.Append(b, "Providers", echo.EchoModule()).(providersBuilder)
}

func (b providersBuilder) WithMongo() providersBuilder {
	return builder.Append(b, "Providers", mongo.MongoDBModule()).(providersBuilder)
}

func (b providersBuilder) WithSentry() providersBuilder {
	return builder.Append(b, "Providers", sentry.SentryModule()).(providersBuilder)
}

func (b providersBuilder) WithConfig() providersBuilder {
	return builder.Append(b, "Providers", config.NewConfig()).(providersBuilder)
}

func (b providersBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var Start = builder.Register(providersBuilder{}, Options{}).(providersBuilder)
