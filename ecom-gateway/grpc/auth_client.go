package grpc

import (
	"fmt"

  "google.golang.org/grpc/credentials/insecure"
  "ecom-gateway/config"
	pbAuth "ecom-gateway/proto/auth"
	"google.golang.org/grpc"
)
var AuthClient pbAuth.AuthServiceClient

func InitAuthClient() {
  fmt.Println("auth url: ", config.GrpcConfig.AuthGRPC)
	conn, err := grpc.Dial(
      config.GrpcConfig.AuthGRPC,
      grpc.WithTransportCredentials(insecure.NewCredentials()),
    )
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	AuthClient = pbAuth.NewAuthServiceClient(conn)
}
