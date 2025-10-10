package handler

import (
	"los-service/internal/app"
	"los-service/internal/domain"
	"los-service/internal/infrastructure/cache"

	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	AppService      *app.ApplicationService
	CustomerService *app.CustomerService
	Cache           *cache.TarantoolCache
}

func NewHandler(appService *app.ApplicationService, customerService *app.CustomerService, cache *cache.TarantoolCache) *Handler {
	return &Handler{
		AppService:      appService,
		CustomerService: customerService,
		Cache:           cache,
	}
}

func (h *Handler) Health(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}

func (h *Handler) SubmitApplication(c *fiber.Ctx) error {
	var app domain.Application
	if err := c.BodyParser(&app); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.AppService.SubmitApplication(&app); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "success"})
}
