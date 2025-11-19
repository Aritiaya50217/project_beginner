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

func (u *contractUsecase) Get(id uint) (*domain.Contract, error) {
	return u.repo.GetByID(id)
}

func (u *contractUsecase) ListByUser(userID uint) ([]*domain.Contract, error) {
	return u.repo.ListByUser(userID)
}

func (u *contractUsecase) Approve(id uint) error {
	contract, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}
	contract.Status = domain.Approved
	return u.repo.Update(contract)
}

func (u *contractUsecase) Reject(id uint) error {
	contract, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}
	contract.Status = domain.Rejected
	return u.repo.Update(contract)
}
