package fx

import (
	"github.com/maiaaraujo5/gostart/monitoring"
	"github.com/maiaaraujo5/gostart/monitoring/sentry"
	"go.uber.org/fx"
)

func SentryModule() fx.Option {
	return fx.Options(
		fx.Provide(
			monitoring.NewConfig,
			sentry.Init,
		),
	)
}
