package port

import (
	"context"
	"smart-stock-trading-platform-stock-service/internal/domain"
	"smart-stock-trading-platform-stock-service/internal/utils"
)

type StockUsecase interface {
	FetchQuote(symbol string) (*utils.StockQuote, error)
	FetchQuotes(symbols []string) ([]*utils.StockQuote, error)
	FetchAllQuotes(exchange string, limit int) ([]*utils.StockQuote, error)
	FetchCompayny(symbol string) (*utils.Company, error)
	AddStockBySymbol(ctx context.Context, symbol string) error
	FindStockByID(ctx context.Context, id int) (*domain.Stock, error)
	DeleteStock(ctx context.Context, id int) error
	GetAllStock(ctx context.Context, offset, limit int) ([]*domain.Stock, int64, error)
	FindStockBySymbol(ctx context.Context, symbol string) (*domain.Stock, error)
}
