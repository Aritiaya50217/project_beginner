package domain

import (
	"time"
)

type Account struct {
	ID        string
	Owner     string
	Balance   float64
	CreatedAt time.Time
}

type AccountRepository interface {
	Create(account *Account) error
	FindByID(id string) (*Account, error)
	Update(account *Account) error
}
