package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", // name of the exchange
		"topic",      // type
		true,         // is the exchange durable?
		false,        // is the exchange auto-deleted?
		false,        // is the exchange internal?
		false,        // no-wait?
		nil,          // arguments?
	)
}
