package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	sentry2 "github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"github.com/maiaaraujo5/gostart/broker"
	"github.com/maiaaraujo5/gostart/log/logger"
	"github.com/maiaaraujo5/gostart/monitoring/sentry"
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

	if c.Sentry {
		err := sentry.Init()
		if err != nil {
			return nil, err
		}
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

func (r *rabbitMQ) Subscribe(queue, exchange string, listener broker.Listener) {
	channel, err := r.connection.Channel()
	if err != nil {
		logger.Error("error to connect in channel")
		return
	}

	err = r.declareQueue(queue, exchange, channel)
	if err != nil {
		logger.Error(fmt.Sprintf("error to declare queue [%s]. error: %s", queue, err))
	}

	messages, err := channel.Consume(queue, "", false, false, false, false, nil)
	if err != nil {
		logger.Error(fmt.Sprintf("error to consume messages from queue [%s]. error: %s", queue, err))
		return
	}

	forever := make(chan bool)
	log.Printf("listening queue %s from rabbitmq", queue)
	go func() {
		for message := range messages {
			span := sentry2.StartSpan(context.Background(), "rabbitmq_consume", sentry2.TransactionName(uuid.NewString()))
			r.handleMessage(listener, message)
			span.Finish()
		}
	}()

	<-forever
}

func (r *rabbitMQ) declareQueue(queueName, exchangeName string, channel *amqp.Channel) error {
	queue, err := channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	if exchangeName != "" {
		err := r.bindQueueInExchange(queue.Name, exchangeName, channel)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *rabbitMQ) bindQueueInExchange(queueName, exchangeName string, channel *amqp.Channel) error {
	err := channel.QueueBind(queueName, "", exchangeName, false, nil)
	if err != nil {
		return err
	}
	return nil
}

func (r *rabbitMQ) notifyError(err error) {
	logger.Error(err.Error())
	sentry2.CaptureException(err)
}

func (r *rabbitMQ) handleMessage(listener broker.Listener, message amqp.Delivery) {
	err := listener(message.Body)
	if err != nil {
		r.notifyError(err)
		if err := message.Nack(false, true); err != nil {
			r.notifyError(err)
		}
	}
	err = message.Ack(true)
	if err != nil {
		logger.Error(err.Error())
		sentry2.CaptureException(err)
	}
}
