package http

import (
	"auth-service/internal/domain"
	"auth-service/internal/usecase"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	usecase *usecase.AuthUsecase
}

func NewAuthHandler(usecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: usecase}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	body := domain.RegisterRequest{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	err := h.usecase.Register(body.Email, body.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "register is successfuly",
	})
}
