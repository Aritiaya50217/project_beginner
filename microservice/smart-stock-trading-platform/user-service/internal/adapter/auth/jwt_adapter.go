package auth

import (
	"context"
	"smart-stock-trading-platform-user-service/internal/port"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTAuth struct {
	secret string
}

func NewJWTAuth(secret string) port.AuthService {
	return &JWTAuth{secret: secret}
}

func (j *JWTAuth) GenerateToken(ctx context.Context, userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secret))
}
