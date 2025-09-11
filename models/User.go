package models

import (
	"time"
)

type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	Phone        string    `json:"phone"`
	Nome         string    `json:"nome"`
	UserType     string    `json:"user_type"`
	CreatedAt    time.Time `json:"created_at"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

type Register struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Nome      string    `json:"nome"`
	Phone     string    `json:"phone"`
	UserType  string    `json:"user_type"`
	CreatedAt time.Time `json:"created_at"`
}

// remover senha

type SanitizedUser struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
	Nome     string `json:"nome"`
}

func Sanitize(u *User) SanitizedUser {
	return SanitizedUser{
		ID:       u.ID,
		Email:    u.Email,
		Nome:     u.Email,
		UserType: u.UserType,
	}
}
