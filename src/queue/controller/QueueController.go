package controller

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"

	"learn_go/rabbit"

)

type MessageQueue struct {
	Message string `validate:"required" json:"message"`
}

func Send(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	rabbitCon := rabbit.Rabbit
	validate := validator.New()

	var messageInput MessageQueue

	err := c.BodyParser(&messageInput)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"httpCode": 400, "status": false, "message": "Invalid!"})
	}

	if errValidate := validate.Struct(messageInput); errValidate != nil {
		return c.Status(400).JSON(fiber.Map{"httpCode": 400, "status": false, "message": "Invalid input!", "error": errValidate.Error()})
	}

	ch, err := rabbitCon.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"TestQueue",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	err = ch.PublishWithContext(
		ctx,
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(messageInput.Message),
		},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"httpCode": 200, "status": true, "message": "OK"})
}
