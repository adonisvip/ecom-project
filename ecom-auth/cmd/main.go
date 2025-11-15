package main

import (
	"context"
	"ecom-auth/config"
	g "ecom-auth/grpc"
	"ecom-auth/pkg/postgres"
	"ecom-auth/pkg/redis"
	pb "ecom-auth/proto/auth"
	"fmt"
	"log"
	"net"

	"github.com/sethvargo/go-envconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()

	var cfg config.Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		log.Fatal(err)
	}

	redis.InitConnection(cfg.Redis)
	_, db, err := postgres.PostgresInitConnection(cfg.Postgres)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	postgres.DB = db

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &g.AuthServer{})
	reflection.Register(s)
	fmt.Println("Auth service running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
