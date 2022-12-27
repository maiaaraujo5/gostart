package fx

import (
	"context"
	"github.com/maiaaraujo5/gostart/database/mongodb"
	"github.com/maiaaraujo5/gostart/database/mongodb/plugins/health"
	"github.com/maiaaraujo5/gostart/log/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

func MongoDBModule() fx.Option {
	return fx.Options(
		fx.Provide(
			func(ctx context.Context) []mongodb.Plugin {
				return []mongodb.Plugin{
					health.NewHealth(ctx),
				}
			},
			func(ctx context.Context, plugins []mongodb.Plugin) (*mongo.Client, error) {
				m, err := mongodb.Connect(ctx, plugins...)
				if err != nil {
					logger.Trace(err.Error())
					return nil, err
				}

				return m.Client, nil
			},
		),
	)
}
