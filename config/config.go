package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string, fallback string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func New(filenames ...string) Config {
	godotenv.Load(filenames...)
	return &configImpl{}
}
