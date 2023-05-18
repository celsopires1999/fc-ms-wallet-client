package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	id := "f71d91a3-8aa6-4fe8-8e14-f66d5acd7b5d"
	account, err := NewAccount(id, 10)
	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, id, account.ID)
	assert.Equal(t, 10.0, account.Balance)
}

func TestCreateAccountWithInvalidId(t *testing.T) {
	account, err := NewAccount("fake-id", 0)
	assert.NotNil(t, err)
	assert.Nil(t, account)
	assert.Equal(t, "invalid UUID length: 7", err.Error())
}
