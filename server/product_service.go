package server

import (
	"context"
	"log"

	"github.com/Sunshine9d/golang-order/db"
	pb "github.com/Sunshine9d/golang-order/gen"
)

type ProductServiceServer struct {
	pb.UnimplementedProductServiceServer
	DB db.ProductDB
}

// CreateProduct handles gRPC requests to create a new product
func (s *ProductServiceServer) CreateProduct(
	ctx context.Context, req *pb.CreateProductRequest,
) (*pb.Product, error) {
	log.Printf("Received request to create product: %s", req.Name)
	product, err := s.DB.CreateProduct(req.Name, req.Price, req.Description, req.Stock)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// GetProductById handles gRPC requests to get a product by ID
func (s *ProductServiceServer) GetProductById(
	ctx context.Context, req *pb.GetProductRequest,
) (*pb.Product, error) {
	log.Printf("Received request for product ID: %d", req.Id)
	product, err := s.DB.GetProductByID(req.Id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// UpdateProduct handles gRPC requests to update a product
func (s *ProductServiceServer) UpdateProduct(
	ctx context.Context, req *pb.UpdateProductRequest,
) (*pb.Product, error) {
	log.Printf("Received request to update product ID: %d", req.Id)
	product, err := s.DB.UpdateProduct(req.Id, req.Name, req.Price, req.Description, req.Stock)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// DeleteProduct handles gRPC requests to delete a product
func (s *ProductServiceServer) DeleteProduct(
	ctx context.Context, req *pb.DeleteProductRequest,
) (*pb.Empty, error) {
	log.Printf("Received request to delete product ID: %d", req.Id)
	err := s.DB.DeleteProduct(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

// GetProducts handles gRPC requests to get a list of products
func (s *ProductServiceServer) GetProducts(
	ctx context.Context, req *pb.GetProductsRequest,
) (*pb.Products, error) {
	log.Printf("Received request to get products")
	products, err := s.DB.GetProducts(int(req.Limit), int(req.Offset), req.Name)
	if err != nil {
		return nil, err
	}
	return &pb.Products{Products: products}, nil
}
