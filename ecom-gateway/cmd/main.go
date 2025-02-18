package main

import (
	"log"

	"ecom-gateway/grpc"
	"ecom-gateway/routers"
)

func main() {

	grpc.InitAuthClient()
	// grpc.InitServiceClient()

  r := routers.SetupRouter()

	log.Fatal(r.Run(":8080"))
}
