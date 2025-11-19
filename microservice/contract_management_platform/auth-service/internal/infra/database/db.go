package database

import (
	"auth-service/internal/domain"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("AUTH_DB_HOST"),
		os.Getenv("AUTH_DB_USER"),
		os.Getenv("AUTH_DB_PASSWORD"),
		os.Getenv("AUTH_DB_NAME"),
		os.Getenv("AUTH_DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect DB: %v", err)
	}

	// Auto migrate
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatalf("failed migrate: %v", err)
	}
	return db
}
