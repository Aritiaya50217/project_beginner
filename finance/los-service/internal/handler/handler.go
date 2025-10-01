package handler

import (
	"log"
	"los-service/internal/app"
	"los-service/internal/domain"
	"los-service/internal/infrastructure/cache"

	"net/http"
	"strconv"

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
	var req domain.Application

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	// create customer entity
	customer := &domain.Customer{
		Name: req.CustomerName,
	}

	// insert customer
	customerID, err := h.CustomerService.CreateCustomer(customer)
	if err != nil {
		log.Println("CreateCustomer error:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create customer")
	}

	// create application entity
	application := &domain.Application{
		CustomerID: customerID,
		Amount:     req.Amount,
		Term:       req.Term,
		Status:     "PENDING",
	}

	// insert application
	applicationID, err := h.AppService.SubmitApplication(application)
	if err != nil {
		log.Println("SubmitApplication error:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to submit application")
	}

	// cache to tarantool
	cacheKey := "app_" + strconv.Itoa(applicationID)
	if err := h.Cache.Set(cacheKey, application); err != nil {
		log.Println("Cache Set error: ", err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"application_id": applicationID,
		"customer_id":    customerID,
		"status":         application.Status,
	})
}
