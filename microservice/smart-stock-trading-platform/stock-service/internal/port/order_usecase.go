package port

import (
	"context"
	"smart-stock-trading-platform-stock-service/internal/domain"
)

type OrderUsecase interface {
	CreateOrder(ctx context.Context, order *domain.Order) error
}
