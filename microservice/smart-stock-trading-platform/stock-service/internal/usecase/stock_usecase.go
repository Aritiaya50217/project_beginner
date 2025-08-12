package usecase

import (
	"context"
	"errors"
	"log"
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

func (u *stockUsecase) FindStockBySymbol(ctx context.Context, symbol string) (*domain.Stock, error) {
	return u.repo.FindBySymbol(ctx, symbol)
}

// AddStockBySymbol จะค้นหา symbol ใน Finnhub และบันทึกลง DB ถ้าเจอ
func (u *stockUsecase) AddStockBySymbol(ctx context.Context, symbol string) error {
	// 1. ดึงข้อมูลหุ้นจาก Finnhub API
	data, err := u.FetchQuote(symbol)
	if err != nil {
		return err
	}
	// 2. หา symbol ที่ต้องการ
	if data == nil {
		return errors.New("symbol not found")
	}

	// 3 check symbol in db
	stock, _ := u.FindStockBySymbol(ctx, data.Symbol)
	if stock != nil {
		log.Printf("The duplicate symbol is (%v)\n.", symbol)
		return errors.New("The duplicate symbol")
	}

	return u.repo.Create(ctx, &domain.Stock{
		Symbol:    stock.Symbol,
		Name:      stock.Symbol,
		LastPrice: stock.LastPrice,
	})
}

func (u *stockUsecase) FindStockByID(ctx context.Context, id int) (*domain.Stock, error) {
	return u.repo.FindStockByID(ctx, id)
}

func (u *stockUsecase) DeleteStock(ctx context.Context, id int) error {
	return u.repo.DeleteStock(ctx, id)
}

func (u *stockUsecase) GetAllStock(ctx context.Context, offset, limit int) ([]*domain.Stock, int64, error) {
	return u.repo.GetAllStock(ctx, offset, limit)
}
