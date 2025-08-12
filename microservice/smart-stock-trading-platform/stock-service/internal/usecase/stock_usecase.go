package usecase

import (
	"context"
	"errors"
	"smart-stock-trading-platform-stock-service/internal/domain"
	"smart-stock-trading-platform-stock-service/internal/port"
	"smart-stock-trading-platform-stock-service/internal/utils"
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
func (u *stockUsecase) FetchQuote(symbol string) (*utils.StockQuote, error) {
	return u.provider.FetchQuote(symbol)
}

func (u *stockUsecase) FetchQuotes(symbols []string) ([]*utils.StockQuote, error) {
	return u.provider.FetchQuotes(symbols)
}

func (u *stockUsecase) FetchAllQuotes(exchange string, limit int) ([]*utils.StockQuote, error) {
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

func (u *stockUsecase) FetchCompayny(symbol string) (*utils.Company, error) {
	symbols, err := u.provider.FetchCompayny(symbol)
	if err != nil {
		return nil, err
	}
	return symbols, nil
}

// AddStockBySymbol จะค้นหา symbol ใน Finnhub และบันทึกลง DB ถ้าเจอ
func (u *stockUsecase) AddStockBySymbol(ctx context.Context, symbol string) error {
	// 1. ดึงข้อมูลหุ้นจาก Finnhub API
	stock, err := u.FetchQuote(symbol)
	if err != nil {
		return err
	}
	// 2. หา symbol ที่ต้องการ
	if stock == nil {
		return errors.New("symbol not found")
	}

	return u.repo.Create(ctx, &domain.Stock{
		Symbol:    stock.Symbol,
		Name:      stock.Symbol,
		LastPrice: stock.Price,
	})
}
