package models

import (
	"database/sql"
)

type User struct {
	ID           int          `json:"id"`
	Email        string       `json:"email"`
	PasswordHash string       `json:"password_hash"`
	CreatedAt    sql.NullTime `json:"created_at"`
}

type FindByEmailResponse struct {
	ID        int          `json:"id"`
	Email     string       `json:"email"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
