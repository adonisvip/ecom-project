package main

import (
	"fmt"
	"log"
	"net"
  "context"
	pb "ecom-auth/proto/auth"
  g "ecom-auth/grpc"
  "ecom-auth/pkg/redis"
  "ecom-auth/pkg/postgres"
  "ecom-auth/config"
	"google.golang.org/grpc"
  "github.com/sethvargo/go-envconfig"
  "google.golang.org/grpc/reflection"
)

func main() {
  ctx := context.Background()

	var cfg config.Config
  if err := envconfig.Process(ctx, &cfg); err != nil {
    log.Fatal(err)
  }

	redis.InitConnection(cfg.Redis)
	postgres.PostgresInitConnection(cfg.Postgres)

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
