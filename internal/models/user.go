package models

import (
	"github.com/google/uuid"
	"time"
)

type Role string

const (
	Employee  Role = "employee"
	Moderator Role = "moderator"
)

type User struct {
	ID           uuid.UUID `db:"id"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	Role         string    `db:"role"`
	CreatedAt    time.Time `db:"created_at"`
}

// CreateUser
// IsValidRole
