package main

import (
	"webapi/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	// Enable CORS
	app.Use(cors.New(
		cors.Config{},
	))
	routes.Setup(app)

	app.Listen(":3000")
}
