package model

import (
	"strconv"
	"time"
)

type Item struct {
	ID        int           `json:"id" gorm:"primaryKey"`
	Name      string        `json:"name"`
	ItemPrice *ItemPrice    `json:"item_price,omitempty"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Logs      []*HistoryLog `json:"log,omitempty" gorm:"polymorphic:Log"`
}

func (i *Item) CreateLog(oldData *Item) {
	var logs []*HistoryLog
	if oldData.Name != i.Name {
		logs = append(logs, &HistoryLog{
			Log: AuthenticatedUser.Name + " update name " + oldData.Name + " to " + i.Name,
		})
	}
	if oldData.ItemPrice.Price != i.ItemPrice.Price {
		logs = append(logs, &HistoryLog{
			Log: AuthenticatedUser.Name + " update price " + strconv.Itoa(oldData.ItemPrice.Price) + " to " + strconv.Itoa(i.ItemPrice.Price),
		})
	}
	if oldData.ItemPrice.Stock != i.ItemPrice.Stock {
		logs = append(logs, &HistoryLog{
			Log: AuthenticatedUser.Name + " update stock " + strconv.Itoa(oldData.ItemPrice.Stock) + " to " + strconv.Itoa(i.ItemPrice.Stock),
		})
	}

	i.Logs = logs
}
