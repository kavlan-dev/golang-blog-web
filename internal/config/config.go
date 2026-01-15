package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	ServerHost        string
	ServerPort        string
	CORSAllowedOrigin string
}

func LoadConfig() *Config {
	host := os.Getenv("SERVER_HOST")
	if host == "" {
		host = "localhost"
		log.Printf("Используется хост по умолчанию: %s", host)
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
		log.Printf("Используется порт по умолчанию: %s", port)
	} else if _, err := strconv.Atoi(port); err != nil {
		log.Fatalf("Некорректный порт: %v", err)
	}

	corsOrigin := os.Getenv("CORS_ALLOWED_ORIGIN")
	if corsOrigin == "" {
		corsOrigin = "*"
		log.Printf("Используется CORS origin по умолчанию: %s", corsOrigin)
	}

	return &Config{
		ServerHost:        host,
		ServerPort:        port,
		CORSAllowedOrigin: corsOrigin,
	}
}

func GetServerAddress(cgf *Config) string {
	return cgf.ServerHost + ":" + cgf.ServerPort
}
