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

func (r *contractRepository) GetByID(id uint) (*domain.Contract, error) {
	var contract domain.Contract
	if err := r.db.First(&contract, id).Error; err != nil {
		return nil, err
	}
	return &contract, nil
}

func (r *contractRepository) ListByUser(userID uint) ([]*domain.Contract, error) {
	var list []*domain.Contract
	if err := r.db.Where("user_id = ? ", userID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *contractRepository) Update(contract *domain.Contract) error {
	return r.db.Save(contract).Error
}
