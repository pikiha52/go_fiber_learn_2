package routes

import (
	"learn_go/src/queue/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupQueueRoutes(router fiber.Router) {
	queue := router.Group("queue")

	queue.Post("/", controller.Send)
}
