package repository

import (
	"errors"
	"smart-stock-trading-platform-stock-service/internal/domain"
	"smart-stock-trading-platform-stock-service/internal/port"

	"gorm.io/gorm"
)

type blockchainRepository struct {
	db *gorm.DB
}

func NewBlockchainRepository(db *gorm.DB) port.BlockchainRepository {
	return &blockchainRepository{db: db}
}

func (r *blockchainRepository) GetLastBlock() (*domain.Block, error) {
	var block domain.Block
	err := r.db.Order("id desc").Offset(0).Limit(10).Find(&block).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// ถ้าไม่มี block แรก ให้สร้าง genesis block (index=0, prevHash="0")
		genesis := domain.NewBlock(0, "GENESIS", "0")
		if err := r.SaveBlock(genesis); err != nil {
			return nil, err
		}
		return &genesis, nil
	}

	if err != nil {
		return nil, err
	}

	return &block, nil
}

func (r *blockchainRepository) SaveBlock(block domain.Block) error {
	return r.db.Create(&block).Error
}
