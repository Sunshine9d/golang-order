package postgres

import (
	"database/sql"
	"fmt"
	"log"

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
