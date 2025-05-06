package main

import (
	"fmt"
	"github.com/aventhis/practice_avito/internal/api"
	"github.com/aventhis/practice_avito/internal/auth"
	"github.com/aventhis/practice_avito/internal/config"
	"github.com/aventhis/practice_avito/internal/storage"
	_ "github.com/google/uuid"
	"log"
	"net/http"
)

//Подключение конфигурации (например, из .env или переменных окружения)
//Подключение к базе данных
//Подключение сервисов (auth, storage и т.п.)
//Настройка API
//Запуск сервера

func main() {
	//Что делает программа первой? — Стартует. Значит, конфиги.
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка чтения конфиг файла: %v", err)

	}

	//Какие сервисы зависят друг от друга? — API зависит от storage и auth.
	// Значит, сначала storage → потом auth → потом API.

	//К чему она подключается? — К БД. Значит, инициализируем storage.
	store, err := storage.NewStorage(cfg.Database.DSN)
	if err != nil {
		log.Fatalf("Не удалось инициализировать хранилище: %v", err)
	}
	if err = store.InitDB(); err != nil {
		log.Fatalf("Не удалось инициализировать базу данных: %v", err)
	}

	// создание auth-сервиса, который будет работать с JWT 🔐
	authService := auth.NewAuthService(cfg.Server.JWTSecret)

	//Что ей нужно, чтобы принимать запросы? — API, значит, настраиваем HTTP-сервер.
	apiServer := api.NewAPI(store, authService)

	// Здесь дальше будет запуск сервера
	router := apiServer.SetupRoutes()

	log.Printf("Сервер запущен на порту :%d", cfg.Server.Port)
	if err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), router); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
