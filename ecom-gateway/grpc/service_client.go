package grpc

import (
	"fmt"
  "google.golang.org/grpc/credentials/insecure"
	pbService "ecom-gateway/proto/service"
	"google.golang.org/grpc"
  "ecom-gateway/config"
)

var ServiceClient pbService.EcomServiceClient

func InitServiceClient() {
	conn, err := grpc.Dial(
    config.GrpcConfig.CoreGRPC,
    grpc.WithTransportCredentials(insecure.NewCredentials()),
    // grpc.WithUnaryInterceptor(apmgrpc.NewUnaryClientInterceptor()),
    // grpc.WithStreamInterceptor(apmgrpc.NewStreamClientInterceptor()),
  )
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	ServiceClient = pbService.NewEcomServiceClient(conn)
}
