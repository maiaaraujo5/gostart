package application

import (
	"context"
	"github.com/maiaaraujo5/gostart/config"
	logger "github.com/maiaaraujo5/gostart/log/instance"
	"go.uber.org/fx"
)

func Run(options Options) error {
	config.Load()
	logger.Load()
	return fx.New(
		fx.Options(options.Providers...),
		fx.Provide(
			context.Background,
		),
		fx.Invoke(start),
	).Start(context.Background())
}

func start(lifecycle fx.Lifecycle, params Params) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				if params.Rest != nil {
					return params.Rest.Start(":8080")
				}

				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		},
	)
}
