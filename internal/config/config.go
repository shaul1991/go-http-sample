package config

import (
	"os"
)

type Config struct {
	Env  string
	Port string
}

func LoadConfig() (*Config, error) {
	env := os.Getenv("ENV")
	port := os.Getenv("PORT")

	return &Config{
		Env:  env,
		Port: port,
	}, nil
}
