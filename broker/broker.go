package broker

type Broker interface {
	SendMessage(exchange, routingKey string, mandatory, immediate bool, message interface{}) error
	Subscribe(queue string, listener Listener)
}
