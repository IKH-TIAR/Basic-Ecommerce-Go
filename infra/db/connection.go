package db

import (
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

func GetConnectionString() string {
	// In a real application, you would typically read these values from environment variables or a configuration file
	// host := "localhost"
	// port := 5432
	// user := "postgres"
	// password := "password"
	// dbname := "ecommerce"
	return "host=localhost port=5432 user=postgres password=dipto811 dbname=ecommerce sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) {
	connStr := GetConnectionString()
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}	
	return db, nil
}

