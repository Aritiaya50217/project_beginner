package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func GenerateJWT(userID uint, email, secret string) (string, time.Time, error) {
	exp := time.Now().Add(10 * time.Minute)

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Printf("Error signing token: %v", err)
		return "", time.Time{}, err
	}

	return tokenStr, exp, nil
}
