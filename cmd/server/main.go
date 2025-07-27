package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	api "github.com/khosbilegt/llama-drover/internal/api/router"
	"github.com/khosbilegt/llama-drover/internal/coordinator"
	"github.com/khosbilegt/llama-drover/internal/db"
)

func init() {
	log.Println("Loading environment variables...")
	dotEnvErr := godotenv.Load()
	if dotEnvErr != nil {
		log.Fatalf("Error loading .env file: %v", dotEnvErr)
	}
}

func main() {
	log.Println("Starting server...")

	port := os.Getenv("PORT")
	mongoURI := os.Getenv("MONGO_URI")

	mongoClient, mongoErr := db.NewMongoClient(mongoURI)
	if mongoErr != nil {
		log.Fatalf("Could not connect to MongoDB: %v", mongoErr)
	}
	defer mongoClient.Disconnect(context.TODO())
	log.Println("Connected to MongoDB")

	coordinator.Init(mongoClient.Database("llama-drover"))

	log.Println("Trying to start the server on port:", port)
	serverErr := http.ListenAndServe(":"+port, api.NewRouter())
	if serverErr != nil {
		log.Fatalf("Could not start server: %v", serverErr)
	}

}
