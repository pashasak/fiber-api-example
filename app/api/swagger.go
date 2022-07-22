package api

import (
	"github.com/gofiber/fiber/v2"

	swagger "github.com/arsmn/fiber-swagger/v2"
)

// SwaggerRoute func for describe group of API Docs routes.
func SwaggerRoute(a *fiber.App) {
	// Create routes group.
	route := a.Group("/docs")

	// Routes for GET method:
	route.Get("*", swagger.HandlerDefault)
}
