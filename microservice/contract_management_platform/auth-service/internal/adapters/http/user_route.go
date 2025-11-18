package http

import (
	"auth-service/internal/adapters/token"
	"auth-service/internal/infra/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App, handler *UserHandler, jwtProvider *token.JWTProvider) {
	api := app.Group("/user")
	api.Post("/login", handler.Login)
	
	api.Use(middleware.JWTMiddleware(jwtProvider))
	api.Get("/profile", handler.Profile)
}
