package main

import (
	"log"
	"net/http"

	api "github.com/khosbilegt/llama-drover/internal/api/router"
)

func main() {
	log.Println("Starting server...")
	err := http.ListenAndServe(":8080", api.NewRouter())
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
