package main

import (
	"booking-system-booking-service/internal/app"
	"booking-system-booking-service/internal/config"
	"booking-system-booking-service/internal/domain"
	"booking-system-booking-service/internal/infrastructure/repository"
	"booking-system-booking-service/internal/interface/handler"
	"booking-system-booking-service/internal/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	// 1. Load config
	cfg := config.LoadConfig()

	// 2. Connect to Azure SQL Server
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=true",
		cfg.SQLServerUser,
		cfg.SQLServerPass,
		cfg.SQLServerHost,
		cfg.SQLServerPort,
		cfg.SQLServerDB,
	)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to Azure SQL Server: %v", err)
	}

	// 3. Auto-migrate models
	if err := db.AutoMigrate(&domain.Booking{}, &domain.Item{}); err != nil {
		log.Fatalf("failed to migrate models: %v", err)
	}

	// 4. Init dependencies
	// booking
	bookingRepo := repository.NewBookingRepository(db)
	bookingUsecase := app.NewBookingUsecase(bookingRepo)
	bookingHandler := handler.NewBookingHandler(bookingUsecase)

	// item
	itemRepo := repository.NewItemRepository(db)
	itemUsecase := app.NewItemUsecase(itemRepo)
	itemHandler := handler.NewItemHandler(itemUsecase)

	// 5. Setup router
	r := gin.Default()
	api := r.Group("/v1/api")
	api.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	// Protected routes
	booking := api.Group("/bookings")
	{
		booking.POST("/", bookingHandler.CreateBooking)
		booking.GET("/:id", bookingHandler.GetByID)
		booking.GET("/", bookingHandler.GetAll)
	}

	item := api.Group("/items")
	{
		item.POST("/", itemHandler.CreateItem)
		item.GET("/:id", itemHandler.GetByID)
		item.GET("/", itemHandler.GetAll)
		item.DELETE("/:id", itemHandler.DeleteItem)
	}

	// 6. Start server
	port := cfg.ServerPort
	if port == "" {
		port = "8082"
	}
	log.Printf("Booking service running at port: %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
