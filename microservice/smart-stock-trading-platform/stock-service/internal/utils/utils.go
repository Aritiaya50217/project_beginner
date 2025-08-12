package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckUserFromToken(c *gin.Context) (userID int) {
	// ดึง userID จาก context (เซ็ตโดย middleware)
	userIDFromToken, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	return userIDFromToken.(int)
}
