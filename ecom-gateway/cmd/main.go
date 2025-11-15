package main

import (
	"log"
	"os"

	"ecom-gateway/config"
	"ecom-gateway/grpc"
	"ecom-gateway/routers"
)

func main() {
	// Load gRPC configuration from environment variables
	config.GrpcConfig.AuthGRPC = os.Getenv("AUTH_SERVICE_GRPC")
	config.GrpcConfig.CoreGRPC = os.Getenv("CORE_SERVICE_GRPC")

	// Initialize gRPC clients
	if err := grpc.InitAuthClient(); err != nil {
		log.Fatalf("Failed to initialize auth client: %v", err)
	}
	// grpc.InitServiceClient()

	r := routers.SetupRouter()

	log.Fatal(r.Run(":8080"))
}
