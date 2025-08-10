package port

import "smart-stock-trading-platform-stock-service/internal/domain"

type StockRepository interface {
	Create(stock domain.Stock) error
}
