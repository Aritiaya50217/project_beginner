package main

import (
	"context"
	"fmt"
	"log"
	application "los-service/internal/app"
	"los-service/internal/handler"
	"los-service/internal/infrastructure/cache"
	"los-service/internal/infrastructure/repository"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Timezone
	loc, _ := time.LoadLocation("Asia/Bangkok")
	t := time.Now().In(loc)
	fmt.Println("Bangkok time:", t)

	// --- DB connection ---
	conn := repository.ConnectPostgres()
	defer conn.Close(context.Background())

	customerRepo := repository.NewPostgresCustomer(conn)
	applicationRepo := repository.NewPostgresApplication(conn)

	// --- Cache connection ---
	tarantoolConn := cache.NewTarantoolCache()

	// --- Service ---
	customerService := application.NewCustomerService(customerRepo)
	applicationService := application.NewApplicationService(applicationRepo)

	// --- Handlers ---
	applicationHandler := handler.NewHandler(applicationService, customerService, tarantoolConn)

	// Routers
	app := fiber.New()
	api := app.Group("/api/v1")
	api.Get("/health", applicationHandler.Health)
	api.Post("/applications", applicationHandler.SubmitApplication)
	// api.Get("/applications/:id", applicationHandler.GetApplicationByID)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("LOS service started on port %s\n", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
