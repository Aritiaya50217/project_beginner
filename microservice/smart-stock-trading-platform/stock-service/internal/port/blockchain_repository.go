package port

import "smart-stock-trading-platform-stock-service/internal/domain"

type BlockchainRepository interface {
	GetLastBlock() (*domain.Block, error)
	SaveBlock(block domain.Block) error
}
