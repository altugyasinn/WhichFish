package routes

import (
	"fiber-mongo-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func FishRoute(app *fiber.App) {
	app.Post("/fish", controllers.CreateFish)
	app.Get("/fish/fishes", controllers.GetFishes)
	app.Get("/fish/:fishId", controllers.GetAFish)

}
