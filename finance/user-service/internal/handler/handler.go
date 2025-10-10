package handler

import (
	"net/http"
	"user-service/internal/app"
	"user-service/internal/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type Handler struct {
	Auth *app.AuthService
}

func NewHandler(auth *app.AuthService) *Handler {
	return &Handler{Auth: auth}
}

func (h *Handler) Login(c *fiber.Ctx) error {
	log.Debug("== Login ==")
	body := domain.User{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := h.Auth.Login(body.Username, body.Password)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{"token": token})
}
