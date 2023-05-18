package getaccount

import (
	"testing"

	"github.com/celsopires1999/fc-ms-wallet-client/internal/entity"
	"github.com/celsopires1999/fc-ms-wallet-client/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAccountUseCase_Execute(t *testing.T) {
	account, _ := entity.NewAccount("f71d91a3-8aa6-4fe8-8e14-f66d5acd7b5d", 10.0)
	accountMock := &mocks.AccountGatewayMock{}
	accountMock.On("FindByID", account.ID).Return(account, nil)

	uc := NewGetAccountUseCase(accountMock)
	inputDTO := GetAccountInputDTO{
		ID: "f71d91a3-8aa6-4fe8-8e14-f66d5acd7b5d",
	}

	_, err := uc.Execute(inputDTO)
	assert.Nil(t, err)
	accountMock.AssertExpectations(t)
	accountMock.AssertNumberOfCalls(t, "FindByID", 1)
}
