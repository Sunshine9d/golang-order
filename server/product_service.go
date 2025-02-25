package server

import (
	"context"
	"log"

	"github.com/Sunshine9d/golang-order/gen"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
}

// GetProduct implements gRPC service
func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.Product, error) {
	log.Println("Received gRPC request for product ID:", req.Id)
	return &pb.Product{
		Id:          req.Id,
		Name:        "Laptop",
		Description: "A high-performance laptop",
		Price:       1200.50,
	}, nil
}
