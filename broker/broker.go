package broker

import "context"

type Broker interface {
	SendMessage(ctx context.Context, message interface{}) error
	Subscribe(queue string, listener Listener)
}
