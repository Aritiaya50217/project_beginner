package domain

import "time"

type Auth struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`               // FK
	User      User      `gorm:"foreignKey:UserID;references:ID" json:"user"` // Relation
	Token     string    `gorm:"type:varchar(255)" json:"token"`
	ExpiredAt time.Time `gorm:"not null" json:"expired_at"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
