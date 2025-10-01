package domain

import "time"

type Application struct {
	ID           int       `json:"id"`
	CustomerID   int       `json:"customer_id"`
	CustomerName string    `json:"customer_name"`
	Amount       float64   `json:"amount"`
	Term         int64     `json:"term"`
	Status       string    `json:"status"` // pending, approved, rejected
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type ApplicationRepository interface {
	Insert(app *Application) (int, error)
}
