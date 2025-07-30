package http

import (
	"banking-hexagonal/internal/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	accountService *application.AccountService
}

func NewHandler(accSvc *application.AccountService) *Handler {
	return &Handler{accountService: accSvc}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	r.POST("/accounts", h.CreateAccount)
	r.GET("/accounts/:id", h.GetAccount)
	r.POST("/accounts/:id/deposit", h.Deposit)
	r.POST("/accounts/:id/withdraw", h.Withdraw)
}

func (h *Handler) CreateAccount(c *gin.Context) {
	var req struct {
		Owner string `json:"owner" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	acc, err := h.accountService.CreateAccount(req.Owner)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, acc)
}

func (h *Handler) GetAccount(c *gin.Context) {
	id := c.Param("id")
	acc, err := h.accountService.GetAccount(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}
	c.JSON(http.StatusOK, acc)
}

func (h *Handler) Deposit(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Amount float64 `json:"amount" binding:"required,gt=0"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.accountService.Deposit(id, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deposit successful"})
}

func (h *Handler) Withdraw(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Amount float64 `json:"amount" binding:"required,gt=0"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.accountService.Withdraw(id, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "withdraw successful"})
}
