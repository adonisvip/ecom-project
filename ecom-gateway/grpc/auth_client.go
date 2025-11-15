package grpc

import (
	"fmt"
	"log"

	"ecom-gateway/config"
	pbAuth "ecom-gateway/proto/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var AuthClient pbAuth.AuthServiceClient

func InitAuthClient() error {
	if config.GrpcConfig.AuthGRPC == "" {
		return fmt.Errorf("AUTH_SERVICE_GRPC environment variable is not set")
	}

	fmt.Println("Connecting to auth service at:", config.GrpcConfig.AuthGRPC)
	conn, err := grpc.Dial(
		config.GrpcConfig.AuthGRPC,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("failed to connect to auth service: %w", err)
	}

	AuthClient = pbAuth.NewAuthServiceClient(conn)
	log.Println("Successfully connected to auth service")
	return nil
}
