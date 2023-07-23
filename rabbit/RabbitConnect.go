package rabbit

import amqp "github.com/rabbitmq/amqp091-go"

var Rabbit *amqp.Connection

func ConnectRabbit() interface{} {
	connectRabbit, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}

	Rabbit = connectRabbit

	return connectRabbit
}
