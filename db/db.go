package db

import (
	"database/sql"
	pb "github.com/Sunshine9d/golang-order/gen"
)

type Database interface {
	Connect() (*sql.DB, error)
}

type ProductDB interface {
	GetProductByID(id int64) (*pb.Product, error)
	CreateProduct(
		name string, price float64, description string, stock int32,
	) (*pb.Product, error)
	UpdateProduct(
		id int64, name string, price float64, description string, stock int32,
	) (*pb.Product, error)
	DeleteProduct(id int64) error
	GetProducts(limit, offset int, name string) ([]*pb.Product, error)
}
