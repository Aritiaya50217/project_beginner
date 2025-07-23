package domain

import "time"

type Account struct {
	ID        string `bson:"_id,omitempty"`
	Owner     string
	Balance   float64
	CreatedAt time.Time
}

type AccountRepository interface {
	Create(account *Account) error
}
