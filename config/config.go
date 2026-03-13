package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    string
}

var config Config

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

	config = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    httpPort,
	}
}

func GetConfig() Config {
	load()
	return config
}
