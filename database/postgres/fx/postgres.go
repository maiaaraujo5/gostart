package fx

import (
	"github.com/maiaaraujo5/gostart/database/postgres"
	"go.uber.org/fx"
)

func PostgresModule() fx.Option {
	return fx.Options(
		fx.Provide(
			postgres.Connect,
		),
	)
}
