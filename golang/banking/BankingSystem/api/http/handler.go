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
