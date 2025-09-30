package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"smart-stock-trading-platform-order-service/internal/adapter/repository"
	"smart-stock-trading-platform-order-service/internal/config"
	"smart-stock-trading-platform-order-service/internal/domain"
	transporthttp "smart-stock-trading-platform-order-service/internal/transport/http"
	"smart-stock-trading-platform-order-service/internal/transport/middleware"
	"smart-stock-trading-platform-order-service/internal/usecase"
	"smart-stock-trading-platform-order-service/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	utils.InitLogger()
	utils.InfoLogger.Println("Order service started...")

	// Load config
	cfg := config.LoadConfig()

	//  Connect to Azure SQL Server
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=true",
		cfg.SQLServerUser,
		cfg.SQLServerPass,
		cfg.SQLServerHost,
		cfg.SQLServerPort,
		cfg.SQLServerDB,
	)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", zap.Error(err))
	}

	if err := db.AutoMigrate(&domain.Order{}); err != nil {
		log.Fatal("failed to migrate database:", zap.Error(err))
	}

	orderRepo := repository.NewOrderRepository(db)
	orderUsecase := usecase.NewOrderUsecase(orderRepo)
	orderHandler := transporthttp.NewOrderHandler(orderUsecase)

	r := gin.Default()
	r.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	api := r.Group("/v1/api/orders")
	api.POST("/", orderHandler.CreateOrder)
	api.GET("/:id", orderHandler.GetOrder)
	api.GET("/users/:user_id/orders", orderHandler.GetOrdersByUser)
	api.DELETE("/:id", orderHandler.DeleteOrder)

	http.Handle("/metrics", promhttp.Handler())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}
	r.Run(":" + port)
}
