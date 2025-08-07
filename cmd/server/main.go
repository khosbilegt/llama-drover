package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/khosbilegt/llama-drover/internal/api/handlers"
	api "github.com/khosbilegt/llama-drover/internal/api/router"
	"github.com/khosbilegt/llama-drover/internal/coordinator"
	"github.com/khosbilegt/llama-drover/internal/db"
	pb "github.com/khosbilegt/llama-drover/internal/model"
	"google.golang.org/grpc"
)

func init() {
	log.Println("Loading environment variables...")

	err := godotenv.Load(".env.server")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	log.Println("Starting server...")
	port := os.Getenv("PORT")
	grpcPort := os.Getenv("GRPC_PORT")
	mongoURI := os.Getenv("MONGO_URI")

	// Default ports if not specified
	if port == "" {
		port = "8080"
	}
	if grpcPort == "" {
		grpcPort = "9090"
	}

	mongoClient, mongoErr := db.NewMongoClient(mongoURI)
	if mongoErr != nil {
		log.Fatalf("Could not connect to MongoDB: %v", mongoErr)
	}
	defer mongoClient.Disconnect(context.TODO())
	log.Println("Connected to MongoDB")

	coordinator.Init(mongoClient.Database("llama-drover"))

	// Use WaitGroup to run both servers concurrently
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		log.Printf("Starting HTTP server on port: %s", port)
		serverErr := http.ListenAndServe(":"+port, api.NewRouter())
		if serverErr != nil {
			log.Fatalf("Could not start HTTP server: %v", serverErr)
		}
	}()

	go func() {
		defer wg.Done()

		listener, err := net.Listen("tcp", ":"+grpcPort)
		if err != nil {
			log.Fatalf("Failed to listen on gRPC port %s: %v", grpcPort, err)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterCoordinatorServer(grpcServer, &handlers.CoordinatorGRPCServer{})

		log.Printf("Starting gRPC server on port: %s", grpcPort)
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Wait for both servers
	wg.Wait()
}
