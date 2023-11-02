package routes

import (
	"github.com/codepnw/blog/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	router := app.Group("/v1")

	router.Get("/", handlers.BlogList)
	router.Post("/", handlers.BlogCreate)
	router.Put("/:id", handlers.BlogUpdate)
	router.Delete("/:id", handlers.BlogDelete)
}
