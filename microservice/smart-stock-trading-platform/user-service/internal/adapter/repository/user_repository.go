package repository

import (
	"context"
	"errors"
	"smart-stock-trading-platform-user-service/internal/domain"
	"smart-stock-trading-platform-user-service/internal/port"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) port.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	return r.db.WithContext(ctx).Create(&user).Error
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).Where("email = ? ", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	tx := r.db.WithContext(ctx).Model(&domain.User{}).
		Where("id = ?", user.ID).
		Updates(domain.User{
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		})

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("no user updated")
	}
	return nil
}
