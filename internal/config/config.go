package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Loggers loggers
	DB      db
	Http    http
}

type loggers struct {
	Level string `yaml:"level"`
}

type db struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

type http struct {
	Server struct {
		HttpServerPort string `yaml:"http_server_port"`
	} `yaml:"server"`
}

func New() *Config {
	configPath := "/Users/mac/Desktop/projekt/redirect_url/config.yaml"
	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("config file does not exiit %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
	return &cfg
}
