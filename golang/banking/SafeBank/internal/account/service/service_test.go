package service_test

import (
	"testing"

	"github.com/Aritiaya50217/project_beginner/internal/account/model"
	"github.com/Aritiaya50217/project_beginner/internal/account/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Create(account model.Account) (*model.Account, error) {
	args := m.Called(account)
	return args.Get(0).(*model.Account), args.Error(1)
}

func (m *MockRepo) FindByID(id int) (*model.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Account), args.Error(1)
}

func (m *MockRepo) Update(account model.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *MockRepo) CreateTransaction(tx model.Transaction) error {
	args := m.Called(tx)
	return args.Error(0)
}

func (m *MockRepo) FindTransactionsByAccountID(id int) ([]model.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).([]model.Transaction), args.Error(1)
}

func TestDeposit_Success(t *testing.T) {
	mockRepo := new(MockRepo)
	s := service.NewAccountService(mockRepo)

	account := &model.Account{ID: 1, Name: "Alice", Balance: 100}
	amount := 50.0

	mockRepo.On("FindByID", 1).Return(account, nil)
	mockRepo.On("Update", mock.Anything).Return(nil)
	mockRepo.On("CreateTransaction", mock.Anything).Return(nil)

	err := s.Deposit(1, amount)
	assert.NoError(t, err)
	assert.Equal(t, 150.0, account.Balance)
	mockRepo.AssertExpectations(t)

}

// กรณี: ฝากเงินติดลบ
func TestDeposit_Invalidamount(t *testing.T) {
	mockRepo := new(MockRepo)
	s := service.NewAccountService(mockRepo)
	err := s.Deposit(1, -100)
	assert.EqualError(t, err, "amount must be positive")
}

// กรณี withdraw ไม่พอเงิน
func TestWithdraw_InsufficientBalance(t *testing.T) {
	mockRepo := new(MockRepo)
	s := service.NewAccountService(mockRepo)

	account := &model.Account{ID: 1, Name: "Bob", Balance: 30}
	mockRepo.On("FindByID", 1).Return(account, nil)

	err := s.Withdraw(1, 100)
	assert.EqualError(t, err, "insufficient balance")
	mockRepo.AssertExpectations(t)
}
