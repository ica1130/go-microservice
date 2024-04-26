package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// we need Consumer in order to consume messages from the queue
type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

// we will be doing that by creating a new consumer and setting it up
func NewConsumer(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}

	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

func (consumer *Consumer) setup() error {
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}

	return declareExchange(channel)
}

type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}
