package entity

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(id string, balance float64) (*Account, error) {
	account := &Account{
		ID:        id,
		Balance:   balance,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := account.validate(); err != nil {
		return nil, err
	}
	return account, nil
}

func (a *Account) validate() error {
	if _, err := uuid.Parse(a.ID); err != nil {
		return err
	}
	return nil
}
