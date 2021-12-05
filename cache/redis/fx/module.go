package fx

import (
	"github.com/maiaaraujo5/gostart/cache/redis"
	"go.uber.org/fx"
)

func RedisModule() fx.Option {
	return fx.Options(
		fx.Provide(
			redis.Connect,
			),
	)
}
