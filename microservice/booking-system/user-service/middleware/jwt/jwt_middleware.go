package jwt

import (
	"booking-system-user-service/internal/app"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func NewAuthMiddleware(authUsecase app.AuthUsecase, secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse JWT
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrInvalidKey
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid jwt"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
			return
		}

		// Manual expiration check (optional)
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
				return
			}
		}

		// Validate token in DB (check revoked + expires_at)
		authToken, err := authUsecase.ValidateToken(context.Background(), tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token revoked or expired"})
			return
		}

		// Set user info to context
		if uid, ok := claims["user_id"].(float64); ok {
			c.Set("user_id", int(uid))
		}
		if email, ok := claims["email"].(string); ok {
			c.Set("email", email)
		}

		c.Set("auth_token", authToken)
		c.Next()

	}
}
