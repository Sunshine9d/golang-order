package main

import (
	"log"
	"net"

	"github.com/Sunshine9d/golang-order/gen"    // Ensure this matches your module name
	"github.com/Sunshine9d/golang-order/server" // Ensure this matches your module name

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	log.Println("Starting gRPC server on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
