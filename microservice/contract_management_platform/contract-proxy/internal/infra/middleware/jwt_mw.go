package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type JWTProvider struct {
	secret []byte
}

func NewJWTProvider() *JWTProvider {
	return &JWTProvider{
		secret: []byte(os.Getenv("JWT_SECRET")),
	}
}

func JWTMiddleware(provider *JWTProvider) fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing or invalid Authorization header",
			})
		}

		tokenStr := strings.TrimPrefix(auth, "Bearer ")

		// Parse Token
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return provider.secret, nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid or expired token",
			})
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token claims",
			})
		}

		// Extract user_id safely
		uidRaw, ok := claims["user_id"]
		if !ok {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "user_id not found in claims",
			})
		}

		uidFloat, ok := uidRaw.(float64)
		if !ok {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid user_id type",
			})
		}

		c.Locals("user_id", uint(uidFloat))
		return c.Next()
	}
}
