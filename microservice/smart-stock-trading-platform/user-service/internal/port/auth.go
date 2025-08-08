package port

import (
	"context"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	GenerateToken(ctx context.Context, userID uint) (string, error)
	ValidateToken(ctx context.Context, token string) (uint, error)
	HashPassword(ctx context.Context, password string) (string, error)
	CheckPasswordHash(ctx context.Context, password, hash string) bool
	GenerateRefreshToken(ctx context.Context, userID uint) (string, error)
	ValidateRefreshToken(ctx context.Context, refreshToken string) (uint, error)
}

type AuthMiddleware interface {
	RequireAuth() gin.HandlerFunc
	RequireInfoUser() gin.HandlerFunc
}
