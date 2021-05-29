package application

import (
	"github.com/maiaaraujo5/gostart/broker"
	"github.com/maiaaraujo5/gostart/rest"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Rest   rest.Rest     `optional:"true"`
	Broker broker.Broker `optional:"true"`
}
