package ports

import "contract-service/internal/domain"

type ContractRepository interface {
	Create(contract *domain.Contract) error
	GetByID(id uint) (*domain.Contract, error)
	ListByUser(userID uint) ([]*domain.Contract, error)
	Update(c *domain.Contract) error
	Delete(id uint) error
}
