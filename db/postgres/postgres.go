package postgres

import (
	"database/sql"
	"fmt"
	"log"

	pb "github.com/Sunshine9d/golang-order/gen"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type PostgresDB struct {
	DB *sql.DB
}

// Connect to PostgreSQL
func NewPostgresDB(connStr string) (*PostgresDB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("PostgreSQL is unreachable: %v", err)
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL")
	return &PostgresDB{DB: db}, nil
}

// GetProductByID fetches a product from PostgreSQL
func (p *PostgresDB) GetProductByID(id int64) (*pb.Product, error) {
	query := "SELECT id, name, price FROM products WHERE id = $1"

	var product pb.Product
	err := p.DB.QueryRow(query, id).Scan(&product.Id, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}

	return &product, nil
}

// GetProductByID fetches a product from PostgreSQL
func (p *PostgresDB) GetProducts() (*pb.Product, error) {
	query := "SELECT id, name, price FROM products WHERE id = $1"

	var product pb.Product
	err := p.DB.QueryRow(query, nil).Scan(&product.Id, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}

	return &product, nil
}
