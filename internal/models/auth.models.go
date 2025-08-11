package models

import "time"

// Satu struct DB untuk tabel users
type UserAuth struct {
	ID         string    `db:"id"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	Role       string    `db:"role_name"`
	CreatedAt  time.Time `db:"created_at"`
	IsVerified bool      `db:"is_verified"`
}

// Satu struct DB untuk tabel profiles
type ProfileAuth struct {
	ID        string    `db:"id"`
	Fullname  string    `db:"fullname"`
	CreatedAt time.Time `db:"created_at"`
}

// Register input body
type RegisterRequest struct {
	Email    string `json:"email" valid:"required"`
	Password string `json:"password" valid:"required"`
	Fullname string `json:"fullname" valid:"required"`
}

// Register output response
type RegisterResponse struct {
	ID       string    `json:"id"`
	Email    string    `json:"email"`
	Fullname string    `json:"fullname"`
	Created  time.Time `json:"created_at"`
}

// Login input body
type LoginRequest struct {
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required"`
}

// Login output data
type LoginResponseData struct {
	ID      string    `json:"id"`
	Email   string    `json:"email"`
	Role    string    `json:"role"`
	Created time.Time `json:"created_at"`
}
