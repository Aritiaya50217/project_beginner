package main

import (
	"context"
	"log"
	"os"
	"time"

	"banking-hexagonal/api/http"
	"banking-hexagonal/internal/application"
	accountRepo "banking-hexagonal/internal/infrastructure/mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Mongo connect error : %v ", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Mongo ping error : %v", err)
	}

	db := client.Database("banking-system")

	accountRepository := accountRepo.NewAccountRepo(db)
	accountService := application.NewAccountService(accountRepository)

	r := gin.Default()
	handler := http.NewHandler(accountService)
	handler.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Gin run error : %v ", err)
	}

}
