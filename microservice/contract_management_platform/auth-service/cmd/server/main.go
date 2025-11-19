package main

import (
	"auth-service/internal/usecase"
	"log"
	"os"

	"auth-service/internal/adapters/http"
	httpHandler "auth-service/internal/adapters/http"
	"auth-service/internal/adapters/repository"
	"auth-service/internal/adapters/token"
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
	authHandler := http.NewAuthHandler(authUsecase)

	userUsecase := usecase.NewUserUsecase(repo)
	userHandler := http.NewUserHandler(userUsecase)

	httpHandler.RegisterAuthRoutes(app, authHandler)
	// middleware
	jwtProvider := token.NewJWTProvider()
	httpHandler.RegisterUserRoutes(app, userHandler, jwtProvider)

	port := os.Getenv("AUTH_PORT")
	if port == "" {
		port = "8081"
	}
	log.Println("auth-service running on port:", port)
	app.Listen(":" + port)
}
