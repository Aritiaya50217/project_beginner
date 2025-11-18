package token

import (
	"auth-service/internal/pkg/utils"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type JWTProvider struct{}

func NewJWTProvider() *JWTProvider {
	return &JWTProvider{}
}

func (j *JWTProvider) ValidateToken(tokenStr string) (map[string]interface{}, error) {
	token, err := utils.ParseToken(tokenStr)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}
	return claims, nil
}
