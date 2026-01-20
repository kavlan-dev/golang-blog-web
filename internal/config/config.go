package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Env   string `json:"env,omitempty"` // Окружение может быть local, dev, prod
	Admin struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
		Email    string `json:"email,omitempty"`
	} `json:"admin"`
	Server struct {
		Host string `json:"host,omitempty"`
		Port uint   `json:"port,omitempty"`
	} `json:"server"`
	CORSAllowedOrigin []string `json:"cors_allowed_origin,omitempty"`
}

func LoadConfig() (*Config, error) {
	pathCmd := flag.String(
		"p",
		"config/config.json",
		"Введите относительный путь до файла конфигурации",
	)

	flag.Parse()

	return configFile(*pathCmd)
}

// Возвращает адрес на котором запускается сервер
func (c *Config) ServerAddress() string {
	port := strconv.Itoa(int(c.Server.Port))
	return c.Server.Host + ":" + port
}

func (c *Config) Cors() string {
	return strings.Join(c.CORSAllowedOrigin, ", ")
}

func configFile(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}

	if cfg.Env != "local" && cfg.Env != "dev" && cfg.Env != "prod" {
		return nil, fmt.Errorf("Окружение %s не найдено", cfg.Env)
	}

	return &cfg, nil
}
