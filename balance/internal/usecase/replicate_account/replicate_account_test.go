package replicateaccount

import (
	"context"
	"testing"

	"github.com/celsopires1999/fc-ms-balance/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestReplicateAccountUseCase_Execute(t *testing.T) {
	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	uc := NewReplicateAccountUseCase(mockUow)
	ctx := context.Background()
	inputDTO := BalanceUpdatedInputDTO{
		AccountIDFrom:        "f71d91a3-8aa6-4fe8-8e14-f66d5acd7b5d",
		BalanceAccountIDFrom: 10.0,
		AccountIDTo:          "2179ace6-4de4-4562-86fa-670178f494a5",
		BalanceAccountIDTo:   20,
	}

	err := uc.Execute(ctx, inputDTO)
	assert.Nil(t, err)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
