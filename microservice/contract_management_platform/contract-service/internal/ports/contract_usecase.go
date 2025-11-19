package ports

import "contract-service/internal/domain"

type ContractUsecase interface {
	Create(userID uint, title, desc string) (*domain.Contract, error)
}
