package config

import "os"

type Config struct {
	DB *DBConfig
}

func New() *Config {
	return &Config{
		DB: &DBConfig{
			Conn: &ConnConfig{
				Host:     os.Getenv("POSTGRES_HOST"),
				Port:     os.Getenv("POSTGRES_PORT"),
				User:     os.Getenv("POSTGRES_USER"),
				Password: os.Getenv("POSTGRES_PASSWORD"),
				DB:       os.Getenv("POSTGRES_DB"),
			},
		},
	}
}
