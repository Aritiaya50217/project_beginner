package main

import (
	"log"
	"os"

	"contract-proxy/internal/adapters/http"
	"contract-proxy/internal/adapters/issuer"
	"contract-proxy/internal/infra/middleware"
	"contract-proxy/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// JWT Provider
	jwtProvider := middleware.NewJWTProvider()

	// HTTP client สำหรับ call contract-service
	issuerClient := issuer.NewHTTPClient(os.Getenv("CONTRACT_SERVICE_URL"))

	// Usecase
	proxyUC := usecase.NewProxyUsecase(issuerClient)

	protected := app.Group("/api", middleware.JWTMiddleware(jwtProvider))
	http.RegisterProxyRoutes(protected, proxyUC)

	port := os.Getenv("CONTRACT_PROXY_PORT")
	if port == "" {
		port = "8083"
	}
	log.Println("contract-proxy-service running on port:", port)
	app.Listen(":" + port)
}
