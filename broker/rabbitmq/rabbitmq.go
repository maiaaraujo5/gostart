package rabbitmq

import (
	"encoding/json"
	"github.com/maiaaraujo5/gostart/broker"
	"github.com/streadway/amqp"
	"log"
)

type rabbitMQ struct {
	connection *amqp.Connection
	config     *Config
}

func Connect() (broker.Broker, error) {
	c, err := NewConfig()
	if err != nil {
		return nil, err
	}

	connection, err := amqp.Dial(c.URL)
	if err != nil {
		return nil, err
	}

	return &rabbitMQ{
		connection: connection,
		config:     c,
	}, nil
}

func (r *rabbitMQ) SendMessage(exchange, routingKey string, mandatory, immediate bool, message interface{}) error {
	channel, err := r.connection.Channel()
	if err != nil {
		return err
	}

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = channel.Publish(
		exchange,
		routingKey,
		mandatory,
		immediate,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	return nil
}

func (r *rabbitMQ) Subscribe(queue string, listener broker.Listener) {
	channel, err := r.connection.Channel()
	if err != nil {
		log.Println("error to connect in channel")
		return
	}

	messages, err := channel.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		log.Printf("error to consume messages from queue [%s]. error: %s", queue, err)
		return
	}

	forever := make(chan bool)
	log.Printf("listening queue %s from rabbitmq", queue)
	go func() {
		for message := range messages {
			err := listener(message.Body)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	<-forever
}
