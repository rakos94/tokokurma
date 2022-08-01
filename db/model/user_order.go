package model

import "time"

type UserOrder struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	OrderID   int       `json:"order_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
