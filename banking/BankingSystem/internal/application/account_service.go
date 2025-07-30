package application

import (
	"banking-hexagonal/internal/domain"
	"errors"
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
		ID:        "", // ให้ repo สร้าง id
		Owner:     owner,
		Balance:   0,
		CreatedAt: time.Now(),
	}
	err := s.repo.Create(acc)
	return acc, err
}

func (s *AccountService) GetAccount(id string) (*domain.Account, error) {
	return s.repo.FindByID(id)
}

func (s *AccountService) Deposit(id string, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	acc, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	acc.Balance += amount
	return s.repo.Update(acc)
}

func (s *AccountService) Withdraw(id string, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	acc, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if acc.Balance < amount {
		return errors.New("insufficient balance")
	}
	acc.Balance -= amount
	return s.repo.Update(acc)
}
