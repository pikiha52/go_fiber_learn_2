package main

import (
	"github.com/gofiber/fiber/v2"

	"learn_go/database"
	"learn_go/rabbit"
	"learn_go/router"

)

func main() {
	app := fiber.New()

	database.ConnectDB()

	rabbit.ConnectRabbit()

	router.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":3000")
}
