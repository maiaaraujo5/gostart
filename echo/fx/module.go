package fx

import (
	"context"
	echoLibrary "github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/gostart/echo"
	"github.com/maiaaraujo5/gostart/echo/plugins/health"
	"github.com/maiaaraujo5/gostart/echo/plugins/sentry"
	"go.uber.org/fx"
)

func EchoModule() fx.Option {
	return fx.Options(
		fx.Provide(
			func(ctx context.Context) []echo.Plugin {
				return []echo.Plugin{
					health.NewHealth(ctx),
					sentry.NewSentry(ctx),
				}
			},
			func(ctx context.Context, plugins []echo.Plugin) (*echoLibrary.Echo, error) {
				e, err := echo.NewEcho(ctx, plugins...)
				if err != nil {
					return nil, err
				}

				return e, nil
			},
		),
	)
}
