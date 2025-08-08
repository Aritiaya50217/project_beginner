package middleware

import (
	"net/http"
	"smart-stock-trading-platform-user-service/internal/port"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService port.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ดึง Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			c.Abort()
			return
		}

		// แยก token ออกจาก "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// ตรวจสอบ token ด้วย authService
		userID, err := authService.ValidateToken(c, tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// ใส่ userID ลง context เพื่อให้ handler ดึงไปใช้ได้
		c.Set("userID", userID)

		c.Next()
	}
}
