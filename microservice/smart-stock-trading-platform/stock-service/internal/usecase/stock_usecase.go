package usecase

import (
	"smart-stock-trading-platform-stock-service/internal/domain"
	"smart-stock-trading-platform-stock-service/internal/port"
)

type stockUsecase struct {
	repo     port.StockRepository
	provider port.MarketDataProvider
	pub      port.EventPublisher
}

func NewStockUsecase(repo port.StockRepository, provider port.MarketDataProvider, pub port.EventPublisher) port.StockUsecase {
	return &stockUsecase{repo: repo, provider: provider, pub: pub}
}

// ดึงข้อมูลหุ้นจากภายนอก (ยังไม่บันทึก DB)
func (u *stockUsecase) FetchQuote(symbol string) (*domain.StockQuote, error) {
	return u.provider.GetQuote(symbol)
}
