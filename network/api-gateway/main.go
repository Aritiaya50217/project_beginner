package main

import (
	"api-gateway/middleware"
	"api-gateway/proxy"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Define backend services
	userServices := &proxy.RoundRobin{Servers: []string{
		"http://localhost:8081",
		"http://localhost:8082",
	}}
	orderServices := &proxy.RoundRobin{Servers: []string{
		"http://localhost:8091",
		"http://localhost:8092",
	}}

	// User routes with JWT
	router.PathPrefix("/user").Handler(middleware.JWTAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		target := userServices.Next()
		proxy.NewSingleHostReverseProxy(target).ServeHTTP(w, r)
	})))

	// Order routes without JWT
	router.PathPrefix("/order").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		target := orderServices.Next()
		proxy.NewSingleHostReverseProxy(target).ServeHTTP(w, r)
	}))

	// Rate Limiter — 5 req ต่อ 1 วินาที
	limiter := middleware.NewRateLimiter(5, time.Second)

	// Apply logging middleware globally
	handler := middleware.LoggingMiddleware(
		middleware.JWTAuth(
			middleware.RateLimitMiddleware(limiter, router),
		),
	)

	log.Println("API Gateway running on :8000")
	if err := http.ListenAndServe(":8000", handler); err != nil {
		log.Fatal(err)
	}
}
