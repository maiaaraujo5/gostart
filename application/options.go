package application

import (
	"github.com/lann/builder"
	rabbitmq "github.com/maiaaraujo5/gostart/broker/rabbitmq/fx"
	fx2 "github.com/maiaaraujo5/gostart/cache/redis/fx"
	mongo "github.com/maiaaraujo5/gostart/database/mongodb/fx"
	postgres "github.com/maiaaraujo5/gostart/database/postgres/fx"
	echo "github.com/maiaaraujo5/gostart/echo/fx"
	sentry "github.com/maiaaraujo5/gostart/monitoring/sentry/fx"
	firebase "github.com/maiaaraujo5/gostart/storage/firebase/fx"
	UDPServer "github.com/maiaaraujo5/gostart/udp/server/fx"
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

func (b providersBuilder) WithPostgres() providersBuilder {
	return builder.Append(b, "Providers", postgres.PostgresModule()).(providersBuilder)
}

func (b providersBuilder) WithRabbitMQ() providersBuilder {
	return builder.Append(b, "Providers", rabbitmq.RabbitMQModule()).(providersBuilder)
}

func (b providersBuilder) WithFirebaseStorage() providersBuilder {
	return builder.Append(b, "Providers", firebase.FirebaseStorageModule()).(providersBuilder)
}

func (b providersBuilder) WithRedis() providersBuilder {
	return builder.Append(b, "Providers", fx2.RedisModule()).(providersBuilder)
}

func (b providersBuilder) WithUDPServer() providersBuilder {
	return builder.Append(b, "Providers", UDPServer.UDPServerModule()).(providersBuilder)
}

func (b providersBuilder) WithCustomProvider(provider fx.Option) providersBuilder {
	return builder.Append(b, "Providers", provider).(providersBuilder)
}

func (b providersBuilder) Build() Options {
	return builder.GetStruct(b).(Options)
}

var Start = builder.Register(providersBuilder{}, Options{}).(providersBuilder)
