package models

import "time"

type User struct {
	ID        string    `json:"id"`
	Login     string    `json:"login"`
	Password  string    `json:"-"` // временно plaintext
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
