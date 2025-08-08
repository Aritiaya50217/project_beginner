package domain

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Email     string    `gorm:"uniqueIndex;size:255"`
	Password  string    `gorm:"size:255"`
	FirstName string    `gorm:"size:255"`
	LastName  string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
