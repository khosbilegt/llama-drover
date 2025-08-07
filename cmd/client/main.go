package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Client application starting...")
	err := godotenv.Load(".env.client")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	ollamaAPIURL := os.Getenv("OLLAMA_API_URL")

	if ollamaAPIURL == "" {
		fmt.Println("OLLAMA_API_URL is not set. Please set it in your environment variables.")
		os.Exit(1)
	}
	fmt.Printf("Connecting to OLLAMA API at %s...\n", ollamaAPIURL)
}
