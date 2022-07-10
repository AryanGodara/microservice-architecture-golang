package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", // name of duration
		"topic",      // type of duration
		true,         // is this exchange durable?
		false,        // get rid of its, when done with it (auto-delete)?
		false,        // is this an exchange just used internally
		false,        // no-wait?
		nil,          // arguments?
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name
		false, // durable?
		false, // delete when unused?
		true,  // is this channel exclusive for current operations?
		false, // no-wait?
		nil,   // arguments?
	)
}
