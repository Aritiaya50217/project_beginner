package service

import (
	"errors"

	"github.com/Aritiaya50217/project_beginner/internal/account/model"
	repository "github.com/Aritiaya50217/project_beginner/internal/account/repository"
)

// AccountService interface กำหนดฟังก์ชันหลักของ account service
type AccountService interface {
	CreateAccount(account model.Account) (*model.Account, error)
	GetAccount(id int) (*model.Account, error)
	Deposit(id int, amount float64) error
	Withdraw(id int, amount float64) error
	GetTransactions(id int) ([]model.Transaction, error)
}

// accountService struct เก็บ repository ไว้ใช้
type accountService struct {
	repo repository.AccountRepository
}

// NewAccountService สร้าง service instance ใหม่
func NewAccountService(repo repository.AccountRepository) AccountService {
	return &accountService{repo: repo}
}

// CreateAccount
func (s *accountService) CreateAccount(account model.Account) (*model.Account, error) {
	account.Balance = 0 // ยอดเงินเริ่มต้น
	return s.repo.Create(account)
}

// GetAccount ดึงบัญชีตาม ID
func (s *accountService) GetAccount(id int) (*model.Account, error) {
	return s.repo.FindByID(id)
}

// Deposit ฝากเงิน
func (s *accountService) Deposit(id int, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	account, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	account.Balance += amount

	// update
	err = s.repo.Update(*account)
	if err != nil {
		return err
	}

	// บันทึกธุรกรรมฝาก
	return s.repo.CreateTransaction(model.Transaction{
		AccountID: id,
		Amount:    amount,
		Type:      "deposit",
	})
}

// Withdraw ถอนเงิน
func (s *accountService) Withdraw(id int, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}

	account, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if account.Balance < amount {
		return errors.New("insufficient balance")
	}

	account.Balance -= amount
	err = s.repo.Update(*account)
	if err != nil {
		return err
	}

	// บันทึกธุรกรรมถอน
	return s.repo.CreateTransaction(model.Transaction{
		AccountID: id,
		Amount:    amount,
		Type:      "withdraw",
	})
}

// GetTransactions ดึงรายการธุรกรรมของบัญชี
func (s *accountService) GetTransactions(id int) ([]model.Transaction, error) {
	return s.repo.FindTransactionByAccountID(id)
}
