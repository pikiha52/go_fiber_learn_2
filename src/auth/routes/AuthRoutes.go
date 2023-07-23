package AuthRoutes

import (
	"github.com/gofiber/fiber/v2"

	"learn_go/src/auth/controller"

)

func SetupAuthRoutes(router fiber.Router) {
	auth := router.Group("auth")

	auth.Post("/signin", controller.Signin)
}
