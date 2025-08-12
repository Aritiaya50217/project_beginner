package repository

import (
	"context"
	"log"
	"smart-stock-trading-platform-stock-service/internal/domain"
	"smart-stock-trading-platform-stock-service/internal/port"

	"gorm.io/gorm"
)

// stockRepository implements port.StockRepository
type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) port.StockRepository {
	return &stockRepository{db: db}
}

func (r *stockRepository) Create(ctx context.Context, stock *domain.Stock) error {
	return r.db.WithContext(ctx).Create(stock).Error
}

func (r *stockRepository) FindBySymbol(ctx context.Context, symbol string) (*domain.Stock, error) {
	var stock domain.Stock
	if err := r.db.WithContext(ctx).Where("symbol = ? ", symbol).First(&stock).Error; err != nil {
		return nil, err
	}
	return &stock, nil
}

func (r *stockRepository) FindStockByID(ctx context.Context, id int) (*domain.Stock, error) {
	var stock domain.Stock
	if err := r.db.WithContext(ctx).First(&stock, id).Error; err != nil {
		log.Printf("error: %+v", err)
		return nil, err
	}
	return &stock, nil
}

func (r *stockRepository) DeleteStock(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&domain.Stock{ID: id}).Error
}
