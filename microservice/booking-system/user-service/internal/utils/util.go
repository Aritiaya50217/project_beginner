package utils

import (
	"booking-system-user-service/internal/domain"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
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

func StartExpiredTokenCleanupJob(db *gorm.DB) {
	go func() {
		for {
			time.Sleep(1 * time.Hour)

			log.Println("[Cron] Cleaning up expired tokens...")

			if err := db.Where("expired_at < ?", time.Now()).Delete(&domain.Auth{}).Error; err != nil {
				log.Printf("[Cron] Failed to delete expired tokens: %v", err)
			}

			log.Println("[Cron] Expired tokens cleaned up")
		}
	}()
}
