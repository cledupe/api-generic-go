package main

import (
	"github.com/cledupe/api-generic-go/database"
	"github.com/cledupe/api-generic-go/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListAllFact)
	app.Post("/fact", handlers.CreatFact)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
