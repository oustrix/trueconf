package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		HTTP  `yaml:"http"`
		Store `yaml:"store"`
	}

	HTTP struct {
		Port string `yaml:"port" env:"HTTP_PORT" env-default:"3333"`
	}

	Store struct {
		Users string `yaml:"users" env:"STORE_USERS" env-default:"users.json"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf(".yml config error: %w", err)
	}

	err = godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("load .env error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf(".env config error: %w", err)
	}

	return cfg, nil
}
