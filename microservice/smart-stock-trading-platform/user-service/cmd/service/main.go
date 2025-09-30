package main

import (
	"fmt"
	"log"
	"net/http"
	"smart-stock-trading-platform-user-service/internal/adapter/auth"
	"smart-stock-trading-platform-user-service/internal/adapter/repository"
	"smart-stock-trading-platform-user-service/internal/config"
	"smart-stock-trading-platform-user-service/internal/domain"
	transporthttp "smart-stock-trading-platform-user-service/internal/transport/http"
	"smart-stock-trading-platform-user-service/internal/transport/middleware"
	"smart-stock-trading-platform-user-service/internal/usecase"

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
		log.Fatal(err)
	}
	// 3. Auto Migrate
	db.AutoMigrate(&domain.User{}, &domain.UserRefreshToken{})

	authService := auth.NewJWTService(cfg.JWTSecret)
	authMiddleware := middleware.AuthMiddleware(authService)

	userRefreshRepository := repository.NewUserRefreshTokenRepository(db)

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, authService, userRefreshRepository)
	userHandler := transporthttp.NewUserHandler(userUsecase, authService)

	r := gin.Default()
	api := r.Group("/v1/api")
	users := api.Group("/users")
	users.POST("/register", userHandler.Register)
	users.POST("/login", userHandler.Login)
	users.Use(authMiddleware)
	{
		users.GET("/:id", userHandler.GetUserByID)
		users.POST("/:id", userHandler.GetUpdateUser)
		users.POST("/refresh-token", userHandler.RefreshToken)
	}

	// 4. Start server
	port := cfg.ServerPort
	if port == "" {
		port = "8081"
	}
	log.Printf("user service running at port: %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
