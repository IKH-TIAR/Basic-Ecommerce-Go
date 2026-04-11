package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(dbConfig config.DatabaseConfig) string {
	// In a real application, you would typically read these values from environment variables or a configuration file
	// host := "localhost"
	// port := 5432
	// user := "postgres"
	// password := "password"
	// dbname := "ecommerce"
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.SSLMode,
	)
	return connString
}

func NewConnection(dbConfig config.DatabaseConfig) (*sqlx.DB, error) {
	connStr := GetConnectionString(dbConfig)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
