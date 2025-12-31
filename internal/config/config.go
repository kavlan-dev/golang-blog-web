package config

import (
	"log"
	"os"
	"strconv"
)

// Config содержит конфигурацию приложения
type Config struct {
	ServerHost        string
	ServerPort        string
	CORSAllowedOrigin string
}

// LoadConfig загружает конфигурацию из переменных окружения или использует значения по умолчанию
func LoadConfig() *Config {
	// Загрузка хоста сервера
	host := os.Getenv("SERVER_HOST")
	if host == "" {
		host = "localhost" // Значение по умолчанию
		log.Printf("Используется хост по умолчанию: %s", host)
	}

	// Загрузка порта сервера
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // Значение по умолчанию
		log.Printf("Используется порт по умолчанию: %s", port)
	} else if _, err := strconv.Atoi(port); err != nil {
		log.Fatalf("Некорректный порт: %v", err)
	}

	// Загрузка разрешенного origin для CORS
	corsOrigin := os.Getenv("CORS_ALLOWED_ORIGIN")
	if corsOrigin == "" {
		corsOrigin = "http://127.0.0.1:5500" // Значение по умолчанию
		log.Printf("Используется CORS origin по умолчанию: %s", corsOrigin)
	}

	return &Config{
		ServerHost:        host,
		ServerPort:        port,
		CORSAllowedOrigin: corsOrigin,
	}
}

// GetServerAddress возвращает полный адрес сервера
func (c *Config) GetServerAddress() string {
	return c.ServerHost + ":" + c.ServerPort
}
