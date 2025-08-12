package domain

import "time"

type Stock struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Symbol      string    `gorm:"uniqueIndex;size:10" json:"symbol"`
	Name        string    `gorm:"size:100" json:"name"`
	CompanyName string    `gorm:"size:100" json:"company_name"`
	LastPrice   float64   `json:"last_price"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}
