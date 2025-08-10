package port

import (
	"smart-stock-trading-platform-stock-service/internal/domain"
)

type MarketDataProvider interface {
	FetchQuote(symbol string) (*domain.StockQuote, error)
	FetchQuotes(symbols []string) ([]*domain.StockQuote, error)
	FetchSymbolList(exchange string) ([]string, error)
}
