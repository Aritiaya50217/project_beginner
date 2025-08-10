package usecase

import (
	"fmt"
	"smart-stock-trading-platform-stock-service/internal/domain"
	"smart-stock-trading-platform-stock-service/internal/port"
)

type stockUsecase struct {
	repo           port.StockRepository
	provider       port.MarketDataProvider
	pub            port.EventPublisher
	blockchainRepo port.BlockchainRepository
}

func NewStockUsecase(repo port.StockRepository, provider port.MarketDataProvider, pub port.EventPublisher, blockchainRepo port.BlockchainRepository) port.StockUsecase {
	return &stockUsecase{repo: repo, provider: provider, pub: pub, blockchainRepo: blockchainRepo}
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

func (u *stockUsecase) AddStockWithBlock(stock domain.Stock) error {
	// หา block ล่าสุด
	lastBlock, err := u.blockchainRepo.GetLastBlock()
	if err != nil {
		return err
	}

	newIndex := lastBlock.Index + 1
	newBlock := domain.NewBlock(newIndex, stock.Symbol, lastBlock.Hash)

	// save block
	if err := u.blockchainRepo.SaveBlock(newBlock); err != nil {
		return err
	}

	// save stock
	stock.BlockHash = newBlock.Hash
	if err := u.repo.Create(stock); err != nil {
		return fmt.Errorf("save stock error: %w", err)
	}
	return nil
}
