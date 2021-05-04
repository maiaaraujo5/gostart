package application

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/maiaaraujo5/gostart/config"
	rest2 "github.com/maiaaraujo5/gostart/rest"
	"go.uber.org/fx"
)

func Run(options Options) error {
	fx.New(options.Providers...)
	return fx.New(
		fx.Provide(
			context.Background,
			config.NewConfig,
		),
		fx.Invoke(start),
	).Start(context.Background())
}

func start(lifecycle fx.Lifecycle, e rest2.Rest) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				log.Infof("starting...")
				e.Start()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Infof("stopping...")
				return nil
			},
		},
	)
}
