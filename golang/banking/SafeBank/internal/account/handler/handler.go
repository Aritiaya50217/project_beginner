package handler

import (
	"net/http"
	"strconv"

	"github.com/Aritiaya50217/project_beginner/internal/account/model"
	"github.com/Aritiaya50217/project_beginner/internal/account/service"
	"github.com/labstack/echo/v4"
)

// AccountHandler โครงสร้าง handler ที่มี service เป็น dependency
type AccountHandler struct {
	service service.AccountService
}

// NewAccountHandler สร้าง handler ใหม่
func NewAccountHandler(s service.AccountService) *AccountHandler {
	return &AccountHandler{service: s}
}

func (h *AccountHandler) CreateAccount(c echo.Context) error {
	req := new(model.Account)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	createAccount, err := h.service.CreateAccount(*req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, createAccount)
}

// GetAccount ดึงข้อมูลบัญชีตาม ID
func (h *AccountHandler) GetAccount(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid account ID"})
	}

	account, err := h.service.GetAccount(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, account)
}

// Deposit ฝากเงิน
func (h *AccountHandler) Deposit(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid account ID"})
	}

	req := struct {
		Amount float64 `json:"amount"`
	}{}

	if err := c.Bind(&req); err != nil || req.Amount <= 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid amount"})
	}

	err = h.service.Deposit(id, req.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "successful"})
}

// Withdraw ถอนเงินจากบัญชี
func (h *AccountHandler) Withdraw(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid account ID"})
	}

	req := struct {
		Amount float64 `json:"amount"`
	}{}

	if err := c.Bind(&req); err != nil || req.Amount <= 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid amount"})
	}

	err = h.service.Withdraw(id, req.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "successful"})
}

// GetTransactions - ดึงประวัติธุรกรรมของบัญชี
func (h *AccountHandler) GetTransactions(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid account ID"})
	}

	transactions, err := h.service.GetTransactions(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, transactions)
}
