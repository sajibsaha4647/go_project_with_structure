package config

import (
	"os"
)

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		return "Not Found"
	}	
	return value
}