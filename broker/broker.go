package broker

import "context"

type Broker interface {
	SendMessage(ctx context.Context, message interface{}) error
	Listen(ctx context.Context, handler Handler)
}
