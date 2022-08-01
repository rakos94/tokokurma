package model

import "time"

type Order struct {
	ID            int            `json:"id" gorm:"primaryKey"`
	CustomerID    int            `json:"customer_id"`
	ItemID        int            `json:"item_id"`
	ItemName      string         `json:"item_name"`
	Quantity      int            `json:"quantity"`
	TotalPrice    int            `json:"total_price"`
	CustomerOrder *CustomerOrder `json:"customer_order,omitempty"`
	UserOrder     *UserOrder     `json:"user_order,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	CustomerName  string         `json:"customer_name"`
	Customer      *Customer      `json:",omitempty"`
}
