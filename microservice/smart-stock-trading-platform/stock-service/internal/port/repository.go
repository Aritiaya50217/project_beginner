package port

import (
	"context"
	"smart-stock-trading-platform-stock-service/internal/domain"
)

type StockRepository interface {
	FindAll(ctx context.Context) ([]domain.Stock, error)
	FindBySymbol(ctx context.Context, symbol string) (*domain.Stock, error)
	Save(ctx context.Context, stock domain.Stock) error
	Create(ctx context.Context, stock *domain.Stock) error
}
