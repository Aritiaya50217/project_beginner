package port

import "smart-stock-trading-platform-stock-service/internal/utils"

type MarketDataProvider interface {
	FetchQuote(symbol string) (*utils.StockQuote, error)
	FetchQuotes(symbols []string) ([]*utils.StockQuote, error)
	FetchSymbolList(exchange string) ([]string, error)
	FetchCompayny(symbol string) (*utils.Company, error)
}
