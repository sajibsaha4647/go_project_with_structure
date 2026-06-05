package config

import (
	"os"
)


type Config struct {
	Port      string
	DBHost    string
	DBUser    string
	DBPass    string
	JWTSecret string
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		return "Not Found"
	}	
	return value
}


func Load() (*Config, error){
	cfg := &Config{
		Port:      GetEnv("PORT",),
		DBHost:    GetEnv("DB_HOST",),
		DBUser:    GetEnv("DB_USER", ),
		DBPass:    GetEnv("DB_PASS", ),
		JWTSecret: GetEnv("JWT_SECRET",),
	}
	return cfg, nil
}