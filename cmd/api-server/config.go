package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// configuring server
func configureServe() *http.Server {
	fmt.Println("Configuring Server...")

	port := os.Getenv("API_PORT")
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      addCorsHeaders(http.DefaultServeMux),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Configuration Complete!")
	return server
}

// adding cors headers
func addCorsHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientUrl := os.Getenv("CLIENT_URL")

		w.Header().Set("Access-Control-Allow-Origin", clientUrl)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
