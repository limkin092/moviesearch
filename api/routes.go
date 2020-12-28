package api

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", HelloWorld)
	app.Get("/BySearch/", BySearch)
}
