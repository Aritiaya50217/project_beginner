package domain

import "time"

type Auth struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`               // FK
	User      User      `gorm:"foreignKey:UserID;references:ID" json:"user"` // Relation
	Token     string    `gorm:"type:varchar(255)" json:"token"`
	Revoked   bool      `gorm:"revoked;default:0" json:"revoked"`
	ExpiredAt time.Time `gorm:"not null" json:"expired_at"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (a *Auth) IsExpired() bool {
	return time.Now().After(a.ExpiredAt)
}
