package handler

import (
	"los-service/internal/app"
	"los-service/internal/domain"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	AppService      *app.ApplicationService
	CustomerService *app.CustomerService
}

func NewHandler(appService *app.ApplicationService, customerService *app.CustomerService) *Handler {
	return &Handler{AppService: appService, CustomerService: customerService}
}

func (h *Handler) Health(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}

func (h *Handler) SubmitApplication(c *fiber.Ctx) error {
	var req domain.Application
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid input",
		})
	}
	req.Status = "pedding"
	if err := h.AppService.SubmitApplication(&req); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	c.Status(http.StatusCreated)
	return c.JSON(fiber.Map{"message": "application submitted"})
}

func (h *Handler) GetApplicationByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}
	result, err := h.AppService.GetApplication(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(result)
}
