package repository

import (
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
