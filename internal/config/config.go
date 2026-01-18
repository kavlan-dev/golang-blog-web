package config

import (
	"encoding/json"
	"flag"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Env   string `json:"env,omitempty"`
	Admin struct {
		Username string `json:"username,omitempty"`
		Password string `json:"password,omitempty"`
		Email    string `json:"email,omitempty"`
	} `json:"admin"`
	Server struct {
		Host string `json:"host,omitempty"`
		Port int    `json:"port,omitempty"`
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

func (c *Config) ServerAddress() string {
	port := strconv.Itoa(c.Server.Port)
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

	return &cfg, nil
}
