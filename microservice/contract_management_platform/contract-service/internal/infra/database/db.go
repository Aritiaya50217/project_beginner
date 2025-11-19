package database

import (
	"contract-service/internal/domain"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("CONTRACT_DB_HOST"),
		os.Getenv("CONTRACT_DB_USER"),
		os.Getenv("CONTRACT_DB_PASSWORD"),
		os.Getenv("CONTRACT_DB_NAME"),
		os.Getenv("CONTRACT_DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect DB: %v", err)
	}

	// AutoMigrate
	if err := db.AutoMigrate(&domain.Contract{}); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}

	return db
}
