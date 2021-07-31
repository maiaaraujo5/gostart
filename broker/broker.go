package broker

type Broker interface {
	SendMessage(exchange, routingKey string, mandatory, immediate bool, message interface{}) error
	Subscribe(queue, exchange string, listener Listener)
}
