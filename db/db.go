package db

import "database/sql"
import "github.com/Sunshine9d/golang-order/gen/pb"

type Database interface {
	Connect() (*sql.DB, error)
}

type ProductDB interface {
	GetProductByID(id int64) (*pb.Product, error)
}
