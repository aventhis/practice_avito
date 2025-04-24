package main

import (
	"github.com/aventhis/practice_avito/internal/storage/postgres"
	_ "github.com/google/uuid"
	"log"
	"os"
	"strconv"
)

//Подключение конфигурации (например, из .env или переменных окружения)
//Подключение к базе данных
//Подключение сервисов (auth, storage и т.п.)
//Настройка API
//Запуск сервера

func main() {
	//Что делает программа первой? — Стартует. Значит, конфиги.
	appPort, err := strconv.Atoi(getEnv("APP_PORT", "8080"))
	if err != nil {
		log.Fatalf("invalid DB_PORT: %v", err)
	}

	dbName := getEnv("DB_NAME", "postgres")                // имя базы
	dbUser := getEnv("DB_USER", "postgres")                // логин
	dbPassword := getEnv("DB_PASSWORD", "postgres")        //пароль
	dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432")) // порт БД
	if err != nil {
		log.Fatalf("invalid DB_PORT: %v", err)
	}
	dbHost := getEnv("DB_HOST", "db") // адрес БД
	dbSSLMode := getEnv("DB_SSL_MODE", "disable")

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Println("[WARN] JWT_SECRET not set, using default dev key")
		jwtSecret = "dev-secret-key"
	}
	//К чему она подключается? — К БД. Значит, инициализируем storage.
	storage, err := postgres.NewStorage()

	//Что ей нужно, чтобы принимать запросы? — API, значит, настраиваем HTTP-сервер.
	//Какие сервисы зависят друг от друга? — API зависит от storage и auth.
	// Значит, сначала storage → потом auth → потом API.

}

// Получить переменную окружения или значение по умолчанию
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
