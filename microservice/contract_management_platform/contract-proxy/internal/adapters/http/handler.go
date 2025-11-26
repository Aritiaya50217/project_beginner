package http

import (
	"contract-proxy/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func RegisterProxyRoutes(router fiber.Router, uc *usecase.ProxyUsecase) {
	router.Get("/contracts", func(c *fiber.Ctx) error {
		userID := c.Locals("user_id").(uint)
		data, err := uc.GetContracts(userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Send(data)
	})
}
