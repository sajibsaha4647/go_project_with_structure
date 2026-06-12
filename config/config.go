package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbName    string
	Port      string
	DBPort    string
	DBHost    string
	DBUser    string
	DBPass    string
	JWTSecret string
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	cfg := &Config{
		Port:      GetEnv("PORT"),
		DBPort:    GetEnv("DB_PORT"),
		DBHost:    GetEnv("DB_HOST"),
		DBUser:    GetEnv("DB_USER"),
		DBPass:    GetEnv("DB_PASSWORD"),
		JWTSecret: GetEnv("JWT_SECRET"),
		DbName:    GetEnv("DB_NAME"),
	}
	return cfg, nil
}