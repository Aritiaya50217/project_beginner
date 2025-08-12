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
	return u.provider.FetchQuote(symbol)
}

func (u *stockUsecase) FetchQuotes(symbols []string) ([]*domain.StockQuote, error) {
	return u.provider.FetchQuotes(symbols)
}

func (u *stockUsecase) FetchAllQuotes(exchange string, limit int) ([]*domain.StockQuote, error) {
	symbols, err := u.provider.FetchSymbolList(exchange)
	if err != nil {
		return nil, err
	}

	if limit > 0 && len(symbols) > limit {
		symbols = symbols[:limit]
	}

	quotes, err := u.provider.FetchQuotes(symbols)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

func (u *stockUsecase) FetchCompayny(symbol string) (*domain.Company, error) {
	symbols, err := u.provider.FetchCompayny(symbol)
	if err != nil {
		return nil, err
	}
	return symbols, nil
}
