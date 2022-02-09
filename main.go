package main

import (
	"os"
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}
	app.Listen(port)
}
