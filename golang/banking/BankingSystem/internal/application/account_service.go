package application

import (
	"banking-hexagonal/internal/domain"
	"time"
)

type AccountService struct {
	repo domain.AccountRepository
}

func NewAccountService(r domain.AccountRepository) *AccountService {
	return &AccountService{repo: r}
}

func (s *AccountService) CreateAccount(owner string) (*domain.Account, error) {
	acc := &domain.Account{
		Owner:     owner,
		Balance:   0,
		CreatedAt: time.Now(),
	}

	err := s.repo.Create(acc)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
