package handler

import (
	"booking-system-user-service/internal/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	usecase app.AuthUsecase
}

func NewAuthHandler(usecase app.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecase: usecase}
}

func (h *AuthHandler) ValidateToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is required"})
		return
	}

	status, err := h.usecase.ValidateToken(c, token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "invalid", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": status}) // เช่น status: "valid"
}
