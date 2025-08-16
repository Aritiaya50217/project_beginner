package domain

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Task      string         `json:"task"`
	Payload   string         `json:"payload"`
	Status    string         `json:"status"`   // queued, in_progress, done
	Progress  int            `json:"progress"` // 0-100
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
