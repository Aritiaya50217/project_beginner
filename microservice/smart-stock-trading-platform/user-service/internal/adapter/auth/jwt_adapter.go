package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type jwtService struct {
	secretKey       string
	refreshDuration time.Duration
}

func NewJWTService(secretKey string) *jwtService {
	return &jwtService{
		secretKey:       secretKey,
		refreshDuration: 24 * time.Hour,
	}
}

type refreshClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func (j *jwtService) GenerateToken(ctx context.Context, userID uint) (string, error) {
	exp := time.Now().Add(j.refreshDuration)
	claims := &refreshClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateToken(ctx context.Context, tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &refreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*refreshClaims); ok && token.Valid {
		return claims.UserID, nil
	}
	return 0, errors.New("invalid claims")

}

func (j *jwtService) HashPassword(ctx context.Context, password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (j *jwtService) CheckPasswordHash(ctx context.Context, password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}

// สร้าง Refresh Token เป็น JWT ที่มี userID และ expire
func (j *jwtService) GenerateRefreshToken(ctx context.Context, userID uint) (string, error) {
	expireTime := time.Now().Add(j.refreshDuration)
	claims := &refreshClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

// ตรวจสอบและ validate Refresh Token คืน userID ถ้า valid
func (j *jwtService) ValidateRefreshToken(ctx context.Context, tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &refreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*refreshClaims); ok && token.Valid {
		return claims.UserID, nil
	}

	return 0, errors.New("invalid refresh token")
}
