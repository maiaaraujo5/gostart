package broker

type (
	Listener func([]byte) error
)

var Listeners = make(map[string]Listener)

func RegisterListener(queue string, handler Listener) {
	Listeners[queue] = handler
}

func GetListeners() map[string]Listener {
	return Listeners
}
