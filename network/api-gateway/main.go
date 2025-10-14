package main

import (
	"api-gateway/middleware"
	"api-gateway/proxy"
	"log"
	"net/http"

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

	// Apply logging middleware globally
	loggedRouter := middleware.Logging(router)

	log.Println("API Gateway running on :8000")
	if err := http.ListenAndServe(":8000", loggedRouter); err != nil {
		log.Fatal(err)
	}
}
