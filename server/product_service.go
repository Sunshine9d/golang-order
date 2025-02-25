package server

import (
	"context"
	"log"

	"github.com/Sunshine9d/golang-order/db"
	"github.com/Sunshine9d/golang-order/gen/pb"
)

type ProductServiceServer struct {
	pb.UnimplementedProductServiceServer
	DB db.ProductDB
}

// GetProductByID handles gRPC requests
func (s *ProductServiceServer) GetProductByID(ctx context.Context, req *pb.GetProductRequest) (*pb.Product, error) {
	log.Printf("Received request for product ID: %d", req.Id)
	product, err := s.DB.GetProductByID(req.Id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
