package http

import (
	"contract-service/internal/ports"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type ContractHandler struct {
	uc ports.ContractUsecase
}

func NewContractHandler(uc ports.ContractUsecase) *ContractHandler {
	return &ContractHandler{uc: uc}
}

type createContractReq struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (h *ContractHandler) Create(c *fiber.Ctx) error {
	var req createContractReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	validate := validator.New()

	if err := validate.Struct(req); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errMsg := fmt.Sprintf("Field '%s' failed on '%s' tag\n", e.Field(), e.Tag())
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": errMsg,
			})
		}
	}

	// check user_id form middleware
	uid := c.Locals("user_id")
	if uid == nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "unauthorized",
		})
	}

	userID, ok := uid.(uint)
	if !ok {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	contract, err := h.uc.Create(userID, req.Title, req.Description)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(contract)
}
