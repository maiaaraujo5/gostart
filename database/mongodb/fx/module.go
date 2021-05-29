package fx

import (
	"github.com/maiaaraujo5/gostart/database/mongodb"
	"go.uber.org/fx"
)

func MongoDBModule() fx.Option {
	return fx.Options(
		fx.Provide(
			mongodb.Connect,
		),
	)
}
