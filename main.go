package main

import (
	"fiber-mongo-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Altug's Fish House",
	})

	routes.FishRoute(app)

	// TODO: Ping database connection
	app.Listen(":3162")
}

// GET /fish/5f9f1b9c0b1c9c0b1c9c0b1c
