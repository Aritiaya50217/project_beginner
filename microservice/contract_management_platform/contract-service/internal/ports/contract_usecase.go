package ports

import "contract-service/internal/domain"

type ContractUsecase interface {
	Create(userID uint, title, desc string) (*domain.Contract, error)
	Get(id uint) (*domain.Contract, error)
	ListByUser(userID uint) ([]*domain.Contract, error)
	Approve(id uint) error
	Reject(id uint) error
	DeleteContract(id uint) error
}
