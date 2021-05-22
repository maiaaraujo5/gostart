package application

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/maiaaraujo5/gostart/config"
	"go.uber.org/fx"
)

func Run(options Options) error {
	return fx.New(
		fx.Options(options.Providers...),
		fx.Provide(
			config.NewConfig,
			context.Background,
		),
		fx.Invoke(start),
	).Start(context.Background())
}

func start(lifecycle fx.Lifecycle, params Params) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				log.Infof("starting...")

				if params.Broker != nil {
					go params.Broker.Listen(ctx, params.BrokerHandler)
				}

				if params.Rest != nil {
					params.Rest.Start()
				}

				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Infof("stopping...")
				return nil
			},
		},
	)
}
