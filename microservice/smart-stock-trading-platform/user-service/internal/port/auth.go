package port

import (
	"context"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	GenerateToken(ctx context.Context, userID uint, email string) (string, error)
	ValidateToken(ctx context.Context, token string) (uint, error)
	HashPassword(ctx context.Context, password string) (string, error)
	CheckPasswordHash(ctx context.Context, password, hash string) bool
}

type AuthMiddleware interface {
	RequireAuth() gin.HandlerFunc
	RequireInfoUser() gin.HandlerFunc
}
