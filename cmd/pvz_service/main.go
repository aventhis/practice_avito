package main

import (
	"github.com/aventhis/practice_avito/internal/config"
	"github.com/aventhis/practice_avito/internal/storage/postgres"
	_ "github.com/google/uuid"
	"log"
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
	//К чему она подключается? — К БД. Значит, инициализируем storage.
	storage, err := postgres.NewStorage(сfg.Database.DSN)
	if err := storage.Ping(); err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	if err := storage.InitDB(); err != nil {
		log.Fatalf("Не удалось инициализировать базу данных: %v", err)
	}
	if err != nil {
		log.Fatalf("Не удалось инициализировать хранилище: %v", err)
	}
	//Что ей нужно, чтобы принимать запросы? — API, значит, настраиваем HTTP-сервер.
	//Какие сервисы зависят друг от друга? — API зависит от storage и auth.
	// Значит, сначала storage → потом auth → потом API.

}
