package main

import (
	"auth-service/internal/usecase"
	"log"
	"os"

	"auth-service/internal/adapters/http"
	httpHandler "auth-service/internal/adapters/http"
	"auth-service/internal/adapters/repository"
	"auth-service/internal/infra/database"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	app := fiber.New()

	// connect db
	db := database.Connect()
	repo := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUsecase(repo)
	handler := http.NewAuthHandler(authUsecase)

	httpHandler.RegisterAuthRoutes(app, handler)

	port := os.Getenv("APP_PORT")
	log.Println("Auth service running on port : ", port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
