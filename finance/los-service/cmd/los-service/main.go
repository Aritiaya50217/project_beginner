package main

import (
	"context"
	"log"
	application "los-service/internal/app"
	"los-service/internal/handler"
	"los-service/internal/infrastructure/repository"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// --- DB connection ---
	conn := repository.ConnectPostgres()
	defer conn.Close(context.Background())
	
	customerRepo := repository.NewPostgresCustomer(conn)
	applicationRepo := repository.NewPostgresApplication(conn)

	// --- Cache connection ---
	// tarantoolConn := cache.Connect()
	// _ = tarantoolConn // สามารถส่งไปใช้ใน service later

	// --- Service ---
	customerService := application.NewCustomerService(customerRepo)
	applicationService := application.NewApplicationService(applicationRepo)

	// --- Handlers ---
	applicationHandler := handler.NewHandler(applicationService, customerService)

	// Routers
	app := fiber.New()
	api := app.Group("/api/v1")
	api.Get("/health", applicationHandler.Health)
	api.Post("/applications", applicationHandler.SubmitApplication)
	api.Get("/applications/:id", applicationHandler.GetApplicationByID)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("LOS service started on port %s\n", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
