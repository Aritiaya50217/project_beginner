package port

import "smart-stock-trading-platform-order-service/internal/domain"

type OrderRepository interface {
	Create(order *domain.Order) error
	FindByID(id uint) (*domain.Order, error)
	Update(order *domain.Order) error
	FindByUserID(userID uint, offset, limit int64) ([]*domain.Order, int64, error)
}
