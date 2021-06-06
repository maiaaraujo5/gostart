package fx

import (
	"github.com/maiaaraujo5/gostart/storage/firebase"
	"go.uber.org/fx"
)

func FirebaseStorageModule() fx.Option {
	return fx.Options(
		fx.Provide(
			firebase.NewStorage,
		),
	)
}
