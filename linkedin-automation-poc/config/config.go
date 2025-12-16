package config

import "os"

type Config struct {
	Email string
	Limit int
}

func Load() Config {
	return Config{
		Email: os.Getenv("EMAIL"),
		Limit: 10,
	}
}
