package fx

import (
	"github.com/maiaaraujo5/gostart/udp/server"
	"go.uber.org/fx"
)

func UDPServerModule() fx.Option {
	return fx.Options(
		fx.Provide(
			server.Server,
			),
		)
}