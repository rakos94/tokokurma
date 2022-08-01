package model

import "time"

type CustomerOrder struct {
	ID           int       `json:"id" gorm:"primaryKey"`
	OrderID      int       `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	CustomerHp   string    `json:"customer_hp"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
