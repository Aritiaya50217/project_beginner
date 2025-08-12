package domain

import "time"

type Order struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	UserID     int       `json:"user_id"`
	StockID    int       `json:"stock_id"`
	Stock      *Stock    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TotalStock int       `json:"total"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
