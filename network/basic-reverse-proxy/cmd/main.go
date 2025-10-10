package main

import (
	"basic-reverse-proxy/internal/proxy"
	"log"
	"net/http"
)

func main() {
	target := "http://localhost:8081"
	proxyHandler := proxy.NewReverseProxy(target)

	server := &http.Server{
		Addr:    ":8080",
		Handler: proxyHandler,
	}

	log.Println("Reverse Proxy started on :8080 ", target)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
