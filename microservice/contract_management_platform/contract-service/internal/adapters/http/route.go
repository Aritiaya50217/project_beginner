package http

import (
	"contract-service/internal/infra/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterContractRoutes(app *fiber.App, h *ContractHandler) {
	r := app.Group("/contracts")
	r.Use(middleware.JWTMiddleware())

	r.Post("/", h.Create)
}
