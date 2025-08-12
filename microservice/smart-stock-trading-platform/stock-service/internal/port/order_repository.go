package port

import (
	"context"
	"smart-stock-trading-platform-stock-service/internal/domain"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *domain.Order) error
}
