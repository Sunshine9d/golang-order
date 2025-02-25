package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Sunshine9d/golang-order/db/postgres" // Import your database implementation
	pb "github.com/Sunshine9d/golang-order/gen"      // Import generated protobuf package
	"github.com/Sunshine9d/golang-order/server"      // Import your service implementation
)

func initDB() *postgres.PostgresDB {
	// Example PostgreSQL connection string
	connStr := "host=localhost port=5432 user=golang_user password=strongpassword dbname=golang_order sslmode=disable"
	db, err := postgres.NewPostgresDB(connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func main() {
	db := initDB()
	defer db.DB.Close() // Close only when main exits
	// Start gRPC server
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	productServiceServer := &server.ProductServiceServer{DB: db}
	// Register ProductService
	pb.RegisterProductServiceServer(grpcServer, productServiceServer) // ✅ Register the service
	// Enable gRPC reflection
	reflection.Register(grpcServer)

	// Run gRPC server in a goroutine
	go func() {
		log.Println("🚀 Starting gRPC server on port 50051...")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("❌ Failed to serve gRPC: %v", err)
		}
	}()

	// Start gRPC-Gateway (REST API)
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err = pb.RegisterProductServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts) // ✅ Fix this
	if err != nil {
		log.Fatalf("❌ Failed to register gRPC-Gateway: %v", err)
	}

	log.Println("🌍 Starting REST API on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("❌ Failed to start REST server: %v", err)
	}
}
