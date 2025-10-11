package main

import (
	"basic-reverse-proxy/internal/proxy"
	"log"
	"net/http"
)

func main() {
	backends := []string{
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
	}
	proxyHandler := proxy.NewLoadBalancedReverseProxy(backends)

	server := &http.Server{
		Addr:    ":8080",
		Handler: proxyHandler,
	}

	log.Println("Reverse Proxy started on :8080 ", backends)
	for i, b := range backends {
		log.Printf("Backend #%d -> %s", i+1, b)
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
