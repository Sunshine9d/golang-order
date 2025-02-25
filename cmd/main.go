package main

import (
	"log"
	"net"

	"github.com/Sunshine9d/golang-order/gen"
	"github.com/Sunshine9d/golang-order/server"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	_, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	_ = grpc.NewServer()
	log.Println("Starting gRPC server on port 50051")
	log.Fatalf("Failed to serve: %v", err)
}
