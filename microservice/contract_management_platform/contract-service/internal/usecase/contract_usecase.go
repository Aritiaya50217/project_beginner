package usecase

import (
	"contract-service/internal/domain"
	"contract-service/internal/ports"
)

type contractUsecase struct {
	repo ports.ContractRepository
}

func NewContractUsecase(repo ports.ContractRepository) ports.ContractUsecase {
	return &contractUsecase{repo: repo}
}

func (u *contractUsecase) Create(userID uint, title, desc string) (*domain.Contract, error) {
	contract := &domain.Contract{
		UserID:      userID,
		Title:       title,
		Description: desc,
		Status:      "pending",
	}
	if err := u.repo.Create(contract); err != nil {
		return nil, err
	}
	return contract, nil
}
