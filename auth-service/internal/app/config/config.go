package config

import (
	"errors"
	"os"
	"strconv"
)

type Config struct {
	DB    *DBConfig
	Token *TokenConfig
}

func New() (*Config, error) {
	accessLifeTime := uint64(600)
	refreshLifeTime := uint64(86400)

	r := os.Getenv("REFRESH_LIFE_TIME")
	a := os.Getenv("ACCESS_LIFE_TIME")

	if r != "" {
		num, err := strconv.ParseUint(r, 10, 64)
		if err != nil {
			return nil, errors.New("Wrong format for the lifetime of a refresh token")
		}
		refreshLifeTime = num
	}
	if a != "" {
		num, err := strconv.ParseUint(a, 10, 64)
		if err != nil {
			return nil, errors.New("Wrong format for the lifetime of a access token")
		}
		accessLifeTime = num
	}
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
		Token: &TokenConfig{
			AccessLifeTime:  accessLifeTime,
			RefreshLifeTime: refreshLifeTime,
		},
	}, nil
}
