package main

import (
	"fmt"
	"log"
	"net"

	"ecom-auth/grpc"
	pb "ecom-auth/proto/auth"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &grpc.AuthServer{})

	fmt.Println("Auth service running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
