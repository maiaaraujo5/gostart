package fx

import (
	"github.com/maiaaraujo5/gostart/broker/rabbitmq"
	"go.uber.org/fx"
)

func RabbitMQModule() fx.Option {
	return fx.Options(
		fx.Provide(
			rabbitmq.Connect,
		),
	)
}
