package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"user-service/internal/app"
	"user-service/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Timezone
	loc, _ := time.LoadLocation("Asia/Bangkok")
	t := time.Now().In(loc)
	fmt.Println("Bangkok time:", t)

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "mysecret"
	}
	authService := app.NewAuthService(secret)
	h := handler.NewHandler(authService)

	app := fiber.New()
	app.Post("/login", h.Login)

	port := os.Getenv("USERSERVICE_PORT")
	if port == "" {
		port = "8081"
	}
	
	log.Printf("user service started on port %s\n", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("cannot start server: %v", err)
	}

}
TODO : check router in kong