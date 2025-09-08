package models

import (
	"time"
)

type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Email        string    `json:"email"`
	PasswordHash string    `json:"passwordhash"`
	Nome         string    `json:"nome"`
	Phone        string    `json:"phone"`
	CreatedAt    time.Time `json:"created_at"`
}
