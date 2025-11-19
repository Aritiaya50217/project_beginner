package main

import (
	httpHandler "contract-service/internal/adapters/http"
	"contract-service/internal/adapters/repository"
	"contract-service/internal/infra/database"
	"contract-service/internal/usecase"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db := database.ConnectDB()

	repo := repository.NewContractRepository(db)
	uc := usecase.NewContractUsecase(repo)
	handler := httpHandler.NewContractHandler(uc)

	// health
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok", "service": "contract-service"})
	})

	httpHandler.RegisterContractRoutes(app, handler)

	port := os.Getenv("CONTRACT_PORT")
	if port == "" {
		port = "8082"
	}
	log.Println("contract-service running on port:", port)
	app.Listen(":" + port)
}
