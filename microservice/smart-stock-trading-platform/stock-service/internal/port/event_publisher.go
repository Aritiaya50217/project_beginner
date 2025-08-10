package port

import "smart-stock-trading-platform-stock-service/internal/domain"

type EventPublisher interface {
	PublishPriceUpdate(stock domain.Stock) error
}
