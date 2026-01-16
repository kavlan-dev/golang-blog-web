package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	AdminUsername     string
	AdminPassword     string
	AdminEmail        string
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

	adminUsername := os.Getenv("USERNAME")
	if adminUsername == "" {
		adminUsername = "admin"
		log.Printf("Используется имя пользователя по умолчанию: %s", adminUsername)
	}

	adminPassword := os.Getenv("PASSWORD")
	if adminPassword == "" {
		adminPassword = "admin"
		log.Printf("Используется пароль по умолчанию: %s", adminPassword)
	}

	adminEmail := os.Getenv("EMAIL")
	if adminEmail == "" {
		adminEmail = "admin"
		log.Printf("Используется почта по умолчанию: %s", adminEmail)
	}

	return &Config{
		AdminUsername:     adminUsername,
		AdminPassword:     adminPassword,
		AdminEmail:        adminEmail,
		ServerHost:        host,
		ServerPort:        port,
		CORSAllowedOrigin: corsOrigin,
	}
}

func GetServerAddress(cgf *Config) string {
	return cgf.ServerHost + ":" + cgf.ServerPort
}
