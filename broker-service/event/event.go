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

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name of the queue
		false, // not durable, get deleted when the server restarts
		false, // delete when unused?
		true,  // exclusive, used by only one connection and the queue will be deleted when that connection closes
		false, // no-wait
		nil,   // no specific arguments for this queue
	)
}
