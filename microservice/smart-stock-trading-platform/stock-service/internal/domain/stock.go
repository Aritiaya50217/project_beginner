package domain

import "time"

type Stock struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Symbol    string    `gorm:"uniqueIndex;size:10" json:"symbol"`
	Name      string    `gorm:"size:100" json:"name"`
	LastPrice float64   `json:"last_price"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// ข้อมูล quote ที่มาจาก API ภายนอก (ไม่จำเป็นต้องบันทึก DB)
type StockQuote struct {
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Open      float64 `json:"open"`
	PrevClose float64 `json:"prevClose"`
}
