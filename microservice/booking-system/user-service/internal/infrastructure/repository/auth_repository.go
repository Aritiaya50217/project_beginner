package repository

import (
	"booking-system-user-service/internal/domain"
	"context"

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

// func (r *authRepository) DeleteExpiredTokens(ctx context.Context) error {
// 	return r.db.WithContext(ctx).Where("expired_at <= ?", time.Now()).
// 		Delete(&domain.Auth{}).Error
// }

func (r *authRepository) FindByToken(ctx context.Context, token string) (*domain.Auth, error) {
	var auth domain.Auth
	err := r.db.WithContext(ctx).Where("token = ?", token).First(&auth).Error
	if err != nil {
		return nil, err
	}
	return &auth, nil
}
