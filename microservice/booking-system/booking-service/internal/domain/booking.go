package domain

import "time"

type Booking struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `json:"user_id"`
	ItemID    int       `json:"item_id"`                       // FK
	Item      Item      `gorm:"foreignKey:ItemID" json:"item"` // Relation
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
}
