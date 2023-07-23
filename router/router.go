package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	AuthRoutes "learn_go/src/auth/routes"
	userRoutes "learn_go/src/user/routes"
	QueueRoutes "learn_go/src/queue/routes"

)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	userRoutes.SetupUserRoutes(api)
	AuthRoutes.SetupAuthRoutes(api)
	QueueRoutes.SetupQueueRoutes(api)
}
