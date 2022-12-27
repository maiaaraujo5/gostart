package application

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Rest *echo.Echo `optional:"true"`
}
