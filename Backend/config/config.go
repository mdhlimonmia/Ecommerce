package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     string
	JwtSecretKey string
}

var configuration Config

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		log.Fatal("VERSION not set in .env file")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Fatal("SERVICE_NAME not set in .env file")
		os.Exit(1)
	}

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		log.Fatal("HTTP_PORT not set in .env file")
		os.Exit(1)
	}

	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET_KEY not set in .env file")
		os.Exit(1)
	}

	configuration = Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     port,
		JwtSecretKey: jwtSecret,
	}
}

func GetConfig() Config {
	loadConfig()
	return configuration
}
