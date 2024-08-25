package config

import (
	"os"
)

type Config struct {
	DBUser     string
	DBName     string
	DBPort     string
	DBPassword string
}

func LoadConfig() *Config {

	return &Config{
		DBUser:     getEnv("DBUSER", "user"),
		DBName:     getEnv("DB_NAME", "db-go"),
		DBPort:     getEnv("DB_PORT", "5433"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {

		return value

	}

	return defaultValue

}
