package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config struct to hold the configuration values
type Config struct {
	Version      string
	ServiceName  string
	HttpPort     string
	JwtSecretKey string
}

var configuration *Config

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

	configuration = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     port,
		JwtSecretKey: jwtSecret,
	}
}

func GetConfig() *Config {
	if configuration == nil {
		loadConfig()
	}
	return configuration
}

// Config fro database connection
type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	SslMode  string
}

var dbConfig *DbConfig

func loadDbConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}
	host := os.Getenv("DB_HOST")
	if host == "" {
		log.Fatal("DB_HOST not set in .env file")
		os.Exit(1)
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		log.Fatal("DB_PORT not set in .env file")
		os.Exit(1)
	}

	dbPort, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		log.Fatal("DB_PORT must be a valid integer")
		os.Exit(1)
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		log.Fatal("DB_USER not set in .env file")
		os.Exit(1)
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		log.Fatal("DB_PASSWORD not set in .env file")
		os.Exit(1)
	}

	name := os.Getenv("DB_NAME")
	if name == "" {
		log.Fatal("DB_NAME not set in .env file")
		os.Exit(1)
	}

	sslMode := os.Getenv("ENABLE_SSL_MODE")
	if sslMode == "true" {
		sslMode = "enable"
	} else {
		sslMode = "disable"
	}

	dbConfig = &DbConfig{
		Host:     host,
		Port:     int(dbPort),
		User:     user,
		Password: password,
		DbName:   name,
		SslMode:  sslMode,
	}
}

func GetDbConfig() *DbConfig {
	if dbConfig == nil {
		loadDbConfig()
	}
	return dbConfig
}
