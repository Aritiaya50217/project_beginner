package http

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(app *fiber.App, handler *AuthHandler) {
	api := app.Group("/auth")
	api.Post("/register", handler.Register)
}
