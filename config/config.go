package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type Config struct {
	Version     string
	ServiceName string
	Database    DatabaseConfig
	HttpPort    string
	Secret      string
}

var config *Config

func load() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Failed To load env")
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")

	if serviceName == "" {
		fmt.Println("Service Name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http Port required")
		os.Exit(1)
	}

	secret := os.Getenv("SECRET")
	if secret == "" {
		fmt.Println("No Secret")
		os.Exit(1)
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("DB Host is required")
		os.Exit(1)
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("DB Port is required")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("DB User is required")
		os.Exit(1)
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("DB Password is required")
		os.Exit(1)
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("DB Name is required")
		os.Exit(1)
	}
	dbSSLMode := os.Getenv("DB_SSL_MODE")
	if dbSSLMode == "" {
		dbSSLMode = "disable"
	}

	config = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpPort,
		Secret:      secret,
		Database: DatabaseConfig{
			Host:     dbHost,
			Port:     dbPort,
			User:     dbUser,
			Password: dbPassword,
			Name:     dbName,
			SSLMode:  dbSSLMode,
		},
	}
}

func GetConfig() *Config {
	if config == nil {
		load()
	}
	return config
}
