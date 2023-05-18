package replicateaccount

import (
	"errors"
	"testing"

	"github.com/celsopires1999/fc-ms-wallet-client/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) UpdateBalance(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestReplicateAccountUseCase_Execute(t *testing.T) {
	accountFrom, _ := entity.NewAccount("f71d91a3-8aa6-4fe8-8e14-f66d5acd7b5d", 10.0)
	accountTo, _ := entity.NewAccount("2179ace6-4de4-4562-86fa-670178f494a5", 20.0)
	notFound := &entity.Account{}
	accountMock := &AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)
	accountMock.On("UpdateBalance", mock.Anything).Return(nil)
	accountMock.On("FindByID", accountFrom.ID).Return(accountFrom, nil)
	accountMock.On("FindByID", accountTo.ID).Return(notFound, errors.New("sql: no rows in result set"))

	uc := NewReplicateAccountUseCase(accountMock)
	inputDTO := BalanceUpdatedInputDTO{
		AccountIDFrom:        "f71d91a3-8aa6-4fe8-8e14-f66d5acd7b5d",
		BalanceAccountIDFrom: 10.0,
		AccountIDTo:          "2179ace6-4de4-4562-86fa-670178f494a5",
		BalanceAccountIDTo:   20,
	}

	err := uc.Execute(inputDTO)
	assert.Nil(t, err)
	accountMock.AssertExpectations(t)
	accountMock.AssertNumberOfCalls(t, "FindByID", 2)
	accountMock.AssertNumberOfCalls(t, "Save", 1)
	accountMock.AssertNumberOfCalls(t, "UpdateBalance", 1)
}
