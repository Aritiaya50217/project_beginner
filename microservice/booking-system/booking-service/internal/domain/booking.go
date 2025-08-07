package domain

import (
	"time"
)

type Booking struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	ItemID    *uint     // ต้องเป็น pointer เพื่อให้ nullable ได้
	Item      *Item     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StartTime time.Time `gorm:"start_time" json:"start_time"`
	EndTime   time.Time `gorm:"end_time" json:"end_time"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
