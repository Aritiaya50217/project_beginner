package usecase

import (
	"context"
	"smart-stock-trading-platform-stock-service/internal/domain"
	"smart-stock-trading-platform-stock-service/internal/port"
)

type OrderUsecase struct {
	repo     port.OrderRepository
	provider port.MarketDataProvider
}

func NewOrderUsecase(repo port.OrderRepository, provider port.MarketDataProvider) port.OrderUsecase {
	return &OrderUsecase{repo: repo, provider: provider}
}

func (u *OrderUsecase) CreateOrder(ctx context.Context, order *domain.Order) error {
	return u.repo.CreateOrder(ctx, order)
}
