package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	userRoutes "learn_go/src/user/routes"
	AuthRoutes "learn_go/src/auth/routes"

)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	userRoutes.SetupUserRoutes(api)
	AuthRoutes.SetupAuthRoutes(api)
}
