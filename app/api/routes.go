package api

import (
	"fiber-api-example/app/api/books"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	books.Routes(v1)
}
