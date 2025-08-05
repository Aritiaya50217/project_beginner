package domain

import "time"

type Item struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"name" json:"name"`
	Description string    `gorm:"description" json:"description"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
