package main

import (
	"booking-system-user-service/internal/app"
	"booking-system-user-service/internal/config"
	"booking-system-user-service/internal/domain"
	"booking-system-user-service/internal/infrastructure/repository"
	"booking-system-user-service/internal/interface/handler"
	"booking-system-user-service/internal/utils"
	"booking-system-user-service/middleware/jwt"

	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	// 1. Load configuration
	cfg := config.LoadConfig()
	config.SetTimeZone()

	// 2. Create DSN string for Azure SQL Server
	dsn := fmt.Sprintf(
		"sqlserver://%s:%s@%s:%s?database=%s&encrypt=true",
		cfg.SQLServerUser,
		cfg.SQLServerPass,
		cfg.SQLServerHost,
		cfg.SQLServerPort,
		cfg.SQLServerDB,
	)

	// 3. Connect to Azure SQL Server
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to Azure SQL Server: %v", err)
	}

	// 4. Auto migrate (if needed)
	if err := db.AutoMigrate(&domain.User{}, &domain.Auth{}); err != nil {
		log.Fatalf("failed to migrate schema: %v", err)
	}

	// 5. Initialize dependencies
	authRepo := repository.NewAuthRepository(db)
	authUsecase := app.NewAuthUsecase(authRepo, cfg.JWTSecret)
	authHandler := handler.NewAuthHandler(authUsecase)

	userRepo := repository.NewUserRepository(db)
	userUsecase := app.NewUserUsecase(userRepo, authUsecase, cfg.JWTSecret)
	userHandler := handler.NewUserHandler(userUsecase)

	//  Cron Job ลบ token
	utils.StartExpiredTokenCleanupJob(db)

	// 6. Setup router
	r := gin.Default()
	api := r.Group("/v1/api")

	// Public routes
	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)

	//  Public routes
	api.POST("/auth/validate", authHandler.ValidateToken) // Endpoint ที่ booking-service จะเรียกเช็ค token

	// Protected routes with middleware
	users := api.Group("/users")
	users.Use(jwt.NewAuthMiddleware(authUsecase, cfg.JWTSecret))
	{
		users.GET("/:id", userHandler.GetUserByID)
		users.POST("/:id", userHandler.UpdateUser)
	}

	// 7. Start server
	port := cfg.ServerPort
	if port == "" {
		port = "8081"
	}
	log.Printf("User service running at port: %s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
