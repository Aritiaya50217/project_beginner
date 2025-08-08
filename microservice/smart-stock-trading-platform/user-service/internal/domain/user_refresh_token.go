package domain

import "time"

type UserRefreshToken struct {
	ID           uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"index"`
	RefreshToken string `gorm:"size:512;uniqueIndex"`
	ExpiresAt    time.Time
	CreatedAt    time.Time
}
