package main

import (
	"github.com/gofiber/fiber/v2"
	"moviesearch/api"
	"moviesearch/repository"
)

func main() {

	app := fiber.New()

	repository.InitDatabase()

	api.SetupRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
