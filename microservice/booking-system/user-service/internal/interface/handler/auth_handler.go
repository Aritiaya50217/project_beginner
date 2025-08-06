package handler

import (
	"booking-system-user-service/internal/app"
	"booking-system-user-service/internal/utils"
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
	var req utils.TokenValidationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token is required"})
		return
	}

	status, err := h.usecase.ValidateToken(c, req.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": status}) // เช่น valid, expired, revoked
}
