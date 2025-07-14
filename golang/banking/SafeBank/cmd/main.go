package main

import (
	"log"
	"os"

	"github.com/Aritiaya50217/project_beginner/config"
	handler "github.com/Aritiaya50217/project_beginner/internal/account/handler"
	repository "github.com/Aritiaya50217/project_beginner/internal/account/repository"
	service "github.com/Aritiaya50217/project_beginner/internal/account/service"
	"github.com/labstack/echo/v4"
)

func main() {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "3306")
	dbUser := getEnv("DB_USER", "root")
	dbPass := getEnv("DB_PASSWORD", "root")
	dbName := getEnv("DB_NAME", "safebank")

	// connect MySQL
	db, err := config.NewMySQLConnection(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	accountRepository := repository.NewMySQLAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)
	accountHandler := handler.NewAccountHandler(accountService)

	// echo
	e := echo.New()

	// routing
	e.GET("/accounts/:id", accountHandler.GetAccount)
	e.POST("/accounts", accountHandler.CreateAccount)
	e.POST("/accounts/:id/deposit", accountHandler.Deposit)
	e.POST("/accounts/:id/withdraw", accountHandler.Withdraw)
	e.GET("/accounts/:id/transactions", accountHandler.GetTransactions)

	// run server
	port := getEnv("PORT", "8080")
	e.Logger.Fatal(e.Start(":" + port))
}

func getEnv(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}
