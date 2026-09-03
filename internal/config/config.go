package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	CONFIG_PATH = "./config/local.yaml"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	Token       string `yaml:"token"`
	StoragePath string `yaml:"storage_path"`
	HTTPServer  `yaml:"http_server"`
	DB          `yaml:"database"`
}

type HTTPServer struct {
	Address    string        `yaml:"address" env-default:"localhost:8083"`
	Timeout    time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type DB struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string `yaml:"db_name"`
}

func MustLoad() *Config {
	configPath := CONFIG_PATH
	if configPath == "" {
		log.Fatal("config path is empty")
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("config file doesn't exists: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("failed to read config: %s", configPath)
	}

	return &cfg
}
