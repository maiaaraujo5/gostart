package fx

import (
	"github.com/maiaaraujo5/gostart/echo"
	"go.uber.org/fx"
)

func EchoModule() fx.Option {
	return fx.Options(
		fx.Provide(
			echo.NewEcho,
		),
	)
}
