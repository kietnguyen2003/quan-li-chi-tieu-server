package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseURL string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load("../../.env.local")
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Port:        getEnv("PORT", "8080"),
		JWTSecret:   getEnv("JWT_SECRET", "defaultsecret"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/quan_li_chi_tieu?sslmode=disable"),
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
