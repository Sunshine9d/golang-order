package db

import "database/sql"
import (
	pb "github.com/Sunshine9d/golang-order/gen"
)

type Database interface {
	Connect() (*sql.DB, error)
}

type ProductDB interface {
	GetProductByID(id int64) (*pb.Product, error)
}
