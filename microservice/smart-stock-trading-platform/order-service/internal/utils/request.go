package utils

import "time"

type OrderRequest struct {
	UserID    int       `json:"user_id"`
	StockID   int       `json:"stock_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
