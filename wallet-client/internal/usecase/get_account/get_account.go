package getaccount

import "github.com/celsopires1999/fc-ms-wallet-client/internal/gateway"

type GetAccountInputDTO struct {
	ID string `json:"account_id"`
}

type GetAccountOutputDTO struct {
	ID      string  `json:"account_id"`
	Balance float64 `json:"account_balance"`
}

type GetAccountUseCase struct {
	AccountGateway gateway.AccountGateway
}

func NewGetAccountUseCase(accountGetway gateway.AccountGateway) *GetAccountUseCase {
	return &GetAccountUseCase{
		AccountGateway: accountGetway,
	}
}

func (uc *GetAccountUseCase) Execute(input GetAccountInputDTO) (*GetAccountOutputDTO, error) {
	account, err := uc.AccountGateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}
	output := &GetAccountOutputDTO{
		ID:      account.ID,
		Balance: account.Balance,
	}

	return output, nil
}
