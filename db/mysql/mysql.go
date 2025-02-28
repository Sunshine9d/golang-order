package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	ConnStr string
}

func (m *MySQLDB) Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", m.ConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("MySQL is unreachable: %v", err)
		return nil, err
	}

	fmt.Println("Connected to MySQL")
	return db, nil
}
