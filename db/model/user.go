package model

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Hp        string    `json:"hp"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     string    `json:"-"`
}

var AuthenticatedUser *User
