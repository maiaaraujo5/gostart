package broker

type (
	Listener func([]byte) error
)

type Options struct {
	Handler  Listener
	Exchange string
}

var Listeners = make(map[string]Options)

func RegisterListener(queue, exchange string, handler Listener) {
	Listeners[queue] = Options{
		Handler:  handler,
		Exchange: exchange,
	}
}

func GetListeners() map[string]Options {
	return Listeners
}
