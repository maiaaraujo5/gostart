package application

import (
	"context"
	"github.com/maiaaraujo5/gostart/broker"
	"github.com/maiaaraujo5/gostart/config"
	"github.com/maiaaraujo5/gostart/log/instance"
	"go.uber.org/fx"
)

func Run(options Options) error {
	config.Load()
	instance.Load()
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
				if params.Broker != nil {
					for queue, handler := range broker.GetListeners() {
						go params.Broker.Subscribe(queue, handler)
					}
				}

				if params.Rest != nil {
					params.Rest.Start()
				}

				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		},
	)
}
