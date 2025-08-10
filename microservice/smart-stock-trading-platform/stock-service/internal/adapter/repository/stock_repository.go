package repository

import (
	"context"
	"errors"
	"smart-stock-trading-platform-stock-service/internal/domain"
	"smart-stock-trading-platform-stock-service/internal/port"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// stockRepository implements port.StockRepository
type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) port.StockRepository {
	return &stockRepository{db: db}
}

func (r *stockRepository) FindAll(ctx context.Context) ([]domain.Stock, error) {
	var stocks []domain.Stock
	if err := r.db.WithContext(ctx).Find(&stocks).Error; err != nil {
		return nil, err
	}

	return stocks, nil
}

func (r *stockRepository) FindBySymbol(ctx context.Context, symbol string) (*domain.Stock, error) {
	var stock domain.Stock
	if err := r.db.WithContext(ctx).Where("symbol = ?", symbol).First(&stock).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &stock, nil
}

// Save บันทึกข้อมูลหุ้น (insert หรือ update)
func (r *stockRepository) Save(ctx context.Context, stock domain.Stock) error {
	return r.db.WithContext(ctx).Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "symbol"}},
			DoUpdates: clause.AssignmentColumns([]string{"last_price", "updated_at"}),
		},
	).Create(&stock).Error
}

func (r *stockRepository) Create(ctx context.Context, stock *domain.Stock) error {
	return r.db.WithContext(ctx).Create(stock).Error
}
