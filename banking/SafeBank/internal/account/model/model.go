package model

import "time"

// Account โครงสร้างข้อมูลบัญชีผู้ใช้
type Account struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

// Transaction โครงสร้างข้อมูลธุรกรรมฝาก - ถอน
type Transaction struct {
	ID        int       `json:"id"`
	AccountID int       `json:"account_id"`
	Amount    float64   `json:"amount"`
	Type      string    `json:"type"` // deposit , withdraw
	CreatedAt time.Time `json:"created_at"`
}
