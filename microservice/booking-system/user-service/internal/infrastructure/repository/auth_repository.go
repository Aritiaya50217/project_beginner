package repository

import (
	"booking-system-user-service/internal/domain"
	"context"
	"errors"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(ctx context.Context, auth domain.Auth) error
	FindByToken(ctx context.Context, token string) (*domain.Auth, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Create(ctx context.Context, auth domain.Auth) error {
	if err := r.db.WithContext(ctx).Create(&auth).Error; err != nil {
		return err
	}
	return nil
}

func (r *authRepository) FindByToken(ctx context.Context, token string) (*domain.Auth, error) {
	var auth domain.Auth
	if err := r.db.WithContext(ctx).Where("token = ?", token).First(&auth).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // ไม่เจอ ถือว่า token ใช้ไม่ได้
		}

		return nil, err
	}

	return &auth, nil
}
