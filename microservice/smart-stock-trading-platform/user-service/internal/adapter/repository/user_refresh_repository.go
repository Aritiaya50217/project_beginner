package repository

import (
	"context"
	"errors"
	"smart-stock-trading-platform-user-service/internal/domain"
	"smart-stock-trading-platform-user-service/internal/port"
	"time"

	"gorm.io/gorm"
)

type userRefreshTokenRepository struct {
	db *gorm.DB
}

func NewUserRefreshTokenRepository(db *gorm.DB) port.UserRefreshTokenRepository {
	return &userRefreshTokenRepository{db: db}
}

// SaveRefreshToken บันทึก refresh token พร้อม expire time
func (r *userRefreshTokenRepository)  SaveRefreshToken(ctx context.Context, userID uint, refreshToken string, expiresAt int64) error {
	userRefreshToken := domain.UserRefreshToken{
		UserID:       userID,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Unix(expiresAt, 0),
		CreatedAt:    time.Now(),
	}

	return r.db.WithContext(ctx).Create(&userRefreshToken).Error
}

// IsRefreshTokenValid ตรวจสอบ refresh token ใน DB ว่ายัง valid หรือหมดอายุ
func (r *userRefreshTokenRepository) IsRefreshTokenValid(ctx context.Context, userID uint, refreshToken string) (bool, error) {
	var refresh domain.UserRefreshToken
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND refresh_token = ? AND expires_at > ?", userID, refreshToken, time.Now()).
		First(&refresh).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return true, err
	}
	return true, nil
}

// RevokeRefreshToken ลบ refresh token เมื่อ logout หรือ revoke
func (r *userRefreshTokenRepository) RevokeRefreshToken(ctx context.Context, userID uint, refreshToken string) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND refresh_token = ?", userID, refreshToken).
		Delete(&domain.UserRefreshToken{}).Error
}
