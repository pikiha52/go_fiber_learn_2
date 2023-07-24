package rabbit

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"

)

var Rabbit *amqp.Connection

func ConnectRabbit() interface{} {
	connectRabbit, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic(err)
	}

	Rabbit = connectRabbit

	return connectRabbit
}

func ReceiveQueue() {
	rabbitCon := Rabbit

	channel, err := rabbitCon.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer channel.Close()

	messages, err := channel.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	for message := range messages {
		log.Printf("Received message:  %s\n", message.Body)
	}

	<-forever
}
