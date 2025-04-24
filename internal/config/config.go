package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type ServerConfig struct {
	Port      int
	Env       string
	JWTSecret string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	SSLMode  string
	DSN      string // Строка подключения к базе данных
}

func LoadConfig() (*Config, error) {
	appPort, err := strconv.Atoi(getEnv("APP_PORT", "8080"))
	if err != nil {
		return nil, fmt.Errorf("invalid APP_PORT: %v", err)
	}

	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432")) // порт БД
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %v", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Println("[WARN] JWT_SECRET not set, using default dev key")
		jwtSecret = "dev-secret-key"
	}

	dbConfig := DatabaseConfig{
		Host:     getEnv("DB_HOST", "db"), // адрес БД
		Port:     dbPort,
		Username: getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		Database: getEnv("DB_NAME", "postgres"),
		SSLMode:  getEnv("DB_SSL_MODE", "disable"),
	}

	dbConfig.DSN = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.Username, dbConfig.Password, dbConfig.Database, dbConfig.SSLMode)

	cfg := &Config{
		Database: dbConfig,
		Server: ServerConfig{
			Port:      appPort,
			Env:       getEnv("APP_ENV", "development"),
			JWTSecret: jwtSecret,
		},
	}

	return cfg, nil
}

// Получить переменную окружения или значение по умолчанию
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
