package domain

import "time"

const (
	Pending   = "pending"   // รอการยืนยัน / รอดำเนินการ
	Confirmed = "confirmed" // การสั่งซื้อได้รับการยืนยันแล้ว
	Executed  = "executed"  // คำสั่งซื้อถูกดำเนินการเรียบร้อย (หุ้นถูกซื้อแล้ว)
	Canceled  = "canceled"
	Rejected  = "rejected"
	Expired   = "expired" // คำสั่งซื้อหมดอายุ (ไม่ถูกดำเนินการทันเวลา)
)

type Order struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	StockID   uint      `json:"stock_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
