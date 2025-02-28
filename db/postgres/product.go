package postgres

import (
	"database/sql"
	"fmt"

	pb "github.com/Sunshine9d/golang-order/gen"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func (p *PostgresDB) CreateProduct(
	name string,
	price float64,
	description string,
	stock int32,
) (*pb.Product, error) {
	query := "INSERT INTO products (name, price, description, stock) " +
		"VALUES ($1, $2, $3, $4) RETURNING id"
	var id int64
	err := p.DB.QueryRow(query, name, price, description, stock).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &pb.Product{
		Id:          id,
		Name:        name,
		Price:       price,
		Description: description,
		Stock:       stock,
	}, nil
}

func (p *PostgresDB) GetProductByID(id int64) (*pb.Product, error) {
	query := "SELECT id, name, price, description, stock FROM products WHERE id = $1"
	var product pb.Product
	err := p.DB.QueryRow(query, id).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Description,
		&product.Stock,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err

	}
	return &product, nil
}

func (p *PostgresDB) UpdateProduct(
	id int64,
	name string,
	price float64,
	description string,
	stock int32,
) (*pb.Product, error) {
	query := "UPDATE products SET name = $1, price = $2, description = $3, " +
		"stock = $4 WHERE id = $5 RETURNING id, name, price, description, stock"
	var product pb.Product
	err := p.DB.QueryRow(query, name, price, description, stock, id).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Description,
		&product.Stock,
	)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *PostgresDB) DeleteProduct(id int64) error {
	query := "DELETE FROM products WHERE id = $1"
	_, err := p.DB.Exec(query, id)
	return err
}

func (p *PostgresDB) GetProducts(limit, offset int, name string) ([]*pb.Product, error) {
	query := "SELECT id, name, price, description, stock FROM " +
		"products WHERE name ILIKE $1 LIMIT $2 OFFSET $3"
	rows, err := p.DB.Query(query, "%"+name+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*pb.Product
	for rows.Next() {
		var product pb.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Price,
			&product.Description, &product.Stock); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}
