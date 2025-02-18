package main

import (
	"fmt"
	"log"
	"net"
	pb "ecom-auth/proto/auth"
  g "ecom-auth/grpc"
  "ecom-auth/pkg/redis"
  "ecom-auth/pkg/postgres"
  "ecom-auth/config"
	"google.golang.org/grpc"
)

func main() {
	redis.InitConnection(config.Redis)
	postgres.PostgresInitConnection(config.Postgres)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &g.AuthServer{})

	fmt.Println("Auth service running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
