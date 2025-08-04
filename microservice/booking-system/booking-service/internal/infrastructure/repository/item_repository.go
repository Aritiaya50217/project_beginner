package repository

import (
	"booking-system-booking-service/internal/domain"
	"context"

	"gorm.io/gorm"
)

type ItemRepository interface {
	Create(ctx context.Context, item *domain.Item) error
	FindByID(ctx context.Context, id int) (*domain.Item, error)
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db}
}

func (r *itemRepository) Create(ctx context.Context, item *domain.Item) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *itemRepository) FindByID(ctx context.Context, id int) (*domain.Item, error) {
	var item domain.Item
	if err := r.db.WithContext(ctx).First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
