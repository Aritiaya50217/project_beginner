package port

import (
	"context"
	"smart-stock-trading-platform-stock-service/internal/domain"
)

type StockRepository interface {
	Create(ctx context.Context, stock *domain.Stock) error
	FindBySymbol(ctx context.Context, symbol string) (*domain.Stock, error)
	FindStockByID(ctx context.Context, id int) (*domain.Stock, error)
	DeleteStock(ctx context.Context, id int) error
	GetAllStock(ctx context.Context, offset, limit int) ([]*domain.Stock, int64, error)
}
