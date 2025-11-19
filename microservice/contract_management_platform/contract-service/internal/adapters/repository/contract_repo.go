package repository

import (
	"contract-service/internal/domain"
	"contract-service/internal/ports"

	"gorm.io/gorm"
)

type contractRepository struct {
	db *gorm.DB
}

func NewContractRepository(db *gorm.DB) ports.ContractRepository {
	return &contractRepository{db: db}
}

func (r *contractRepository) Create(contract *domain.Contract) error {
	return r.db.Create(contract).Error
}
