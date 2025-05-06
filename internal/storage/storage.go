package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(DSN string) (*Storage, error) {
	db, err := sql.Open("postgres", DSN)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("не удалось подключиться к базе данных: %w", err)
	}
	return &Storage{db: db}, nil
}

func (s *Storage) InitDB() error {
	query := `
    CREATE TABLE IF NOT EXISTS users(
        id UUID PRIMARY KEY,
        email TEXT NOT NULL UNIQUE,
        password_hash TEXT NOT NULL,
        role TEXT NOT NULL CHECK (role IN ('employee', 'moderator')),
        created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
    );`

	_, err := s.db.Exec(query)
	if err != nil {
		return fmt.Errorf("ошибка создания таблицы users: %w", err)
	}

	return nil
}
