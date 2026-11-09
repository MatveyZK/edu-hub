package entities

import "time"

type User struct {
	ID           int       `json:"id" db:"id"`
	FirstName    string    `json:"first_name" db:"first_name"`
	LastName     string    `json:"last_name" db:"last_name"`
	Email        string    `json:"email" db:"email"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	PasswordHash string    `json:"passsword_hash" db:"passsword_hash"`
	RoleID       int       `json:"role_id" db:"role_id"`
}

type CreateUserEntry struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"passsword_hash"`
	RoleID       int    `json:"role_id"`
}

type UpdateUserEntry struct {
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	Email        *string `json:"email"`
	PasswordHash *string `json:"passsword_hash"`
}
