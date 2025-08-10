package port

import (
	"smart-stock-trading-platform-stock-service/internal/domain"
)

type StockUsecase interface {
	FetchQuote(symbol string) (*domain.StockQuote, error)
	FetchQuotes(symbols []string) ([]*domain.StockQuote, error)
}
