package main

import (
	"fmt"
	"log"
	"net/http"
	"smart-stock-trading-platform-stock-service/internal/adapter/marketdata"
	"smart-stock-trading-platform-stock-service/internal/adapter/publisher"
	"smart-stock-trading-platform-stock-service/internal/adapter/repository"
	"smart-stock-trading-platform-stock-service/internal/config"
	"smart-stock-trading-platform-stock-service/internal/domain"
	transport_http "smart-stock-trading-platform-stock-service/internal/transport/http"
	"smart-stock-trading-platform-stock-service/internal/transport/middleware"
	"smart-stock-trading-platform-stock-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {

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

	if err := db.AutoMigrate(&domain.Stock{}); err != nil {
		log.Fatal("failed to migrate database:", zap.Error(err))
	}

	// Init adapter
	stockRepo := repository.NewStockRepository(db)
	marketProvider := marketdata.NewFinnhubAdapter()
	kafkaPub := publisher.NewKafkaPublisher(cfg.KafkaBrokers)

	// Init usecase
	stockUsecase := usecase.NewStockUsecase(stockRepo, marketProvider, kafkaPub)

	// Handler
	stockHandler := transport_http.NewStockHandler(stockUsecase)

	// Init gin
	r := gin.Default()
	r.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	api := r.Group("/v1/api")
	stock := api.Group("/stocks")
	{
		stock.GET("/finnhub/:symbol", stockHandler.GetQuote)
		stock.GET("/finnhub/", stockHandler.GetAllQuote)
		stock.GET("/finnhub/company/:symbol", stockHandler.GetCompany)
		stock.POST("/:symbol", stockHandler.AddStock)
		stock.GET("/:id", stockHandler.GetStockByID)
		stock.DELETE("/:id", stockHandler.DeleteStock)
		stock.GET("/", stockHandler.GetAllStock)
	}

	// Start server
	port := cfg.ServerPort
	if port == "" {
		port = "8082"
	}

	log.Printf("stock service running at port: %s\n", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("failed to run server: %v", zap.Error(err))
	}

}
