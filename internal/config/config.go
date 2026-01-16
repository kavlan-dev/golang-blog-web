package config

import (
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

func LoadConfig() (*Config, error) {
	host := os.Getenv("SERVER_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	} else if _, err := strconv.Atoi(port); err != nil {
		return nil, err
	}

	corsOrigin := os.Getenv("CORS_ALLOWED_ORIGIN")
	if corsOrigin == "" {
		corsOrigin = "*"
	}

	adminUsername := os.Getenv("ADMIN_USERNAME")
	if adminUsername == "" {
		adminUsername = "admin"
	}

	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminPassword == "" {
		adminPassword = "admin"
	}

	adminEmail := os.Getenv("ADMIN_EMAIL")
	if adminEmail == "" {
		adminEmail = "admin"
	}

	return &Config{
		AdminUsername:     adminUsername,
		AdminPassword:     adminPassword,
		AdminEmail:        adminEmail,
		ServerHost:        host,
		ServerPort:        port,
		CORSAllowedOrigin: corsOrigin,
	}, nil
}

func GetServerAddress(cgf *Config) string {
	return cgf.ServerHost + ":" + cgf.ServerPort
}
