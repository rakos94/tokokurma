package model

import (
	"time"
)

type ItemPrice struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	ItemID    int       `json:"item_id"`
	Price     int       `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
