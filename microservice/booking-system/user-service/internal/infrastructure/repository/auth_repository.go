package repository

import (
	"booking-system-user-service/internal/domain"
	"context"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(ctx context.Context, auth domain.Auth) error
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
