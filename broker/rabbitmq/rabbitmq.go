package rabbitmq

import (
	"context"
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

func (r *rabbitMQ) SendMessage(ctx context.Context, message interface{}) error {
	channel, err := r.connection.Channel()
	if err != nil {
		return err
	}

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	for _, c := range r.config.Sender {
		err = channel.Publish(
			c.Exchange,
			c.RoutingKey,
			c.Mandatory,
			c.Immediate,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
			})

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *rabbitMQ) Listen(ctx context.Context, handler broker.Handler) {
	channel, err := r.connection.Channel()
	if err != nil {
		log.Println("error to connect in channel")
		return
	}

	messages, err := channel.Consume(r.config.Listener.Queue, "", true, false, false, false, nil)
	if err != nil {
		log.Println("error to consume message from queue")
		return
	}

	forever := make(chan bool)

	go func() {
		for message := range messages {
			err := handler.Handle(message.Body)
			if err != nil {
				log.Println(err)
			}
		}
	}()

	<-forever
}
