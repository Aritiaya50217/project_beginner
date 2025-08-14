package port

import "smart-stock-trading-platform-order-service/internal/domain"

type EventPublisher interface {
	PublishOrderCreated(order *domain.Order) error
	PublishOrderUpdated(order *domain.Order) error
}
