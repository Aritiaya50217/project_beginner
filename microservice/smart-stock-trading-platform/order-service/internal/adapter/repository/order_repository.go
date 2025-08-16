package repository

import (
	"smart-stock-trading-platform-order-service/internal/domain"
	"smart-stock-trading-platform-order-service/internal/port"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) port.OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *domain.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) FindByID(id uint) (*domain.Order, error) {
	var order domain.Order
	err := r.db.First(&order, id).Error
	return &order, err
}

func (r *OrderRepository) Update(order *domain.Order) error {
	return r.db.Save(order).Error
}

func (r *OrderRepository) FindByUserID(userID uint, offset, limit int64) ([]*domain.Order, int64, error) {
	var orders []*domain.Order
	var total int64
	// Query สำหรับ total count
	if err := r.db.Model(&domain.Order{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Where("user_id = ?", userID).
		Offset(int(offset)).
		Limit(int(limit)).
		Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	return orders, total, nil
}
