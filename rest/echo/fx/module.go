package fx

import (
	echoLibrary "github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/gostart/rest/echo"
	"go.uber.org/fx"
)

func EchoModule() fx.Option {
	return fx.Options(
		fx.Provide(
			echoLibrary.New,
			echo.NewEcho,
		),
	)
}
