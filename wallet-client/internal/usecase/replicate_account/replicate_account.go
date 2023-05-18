package replicateaccount

import (
	"github.com/celsopires1999/fc-ms-wallet-client/internal/entity"
	"github.com/celsopires1999/fc-ms-wallet-client/internal/gateway"
)

type BalanceUpdatedInputDTO struct {
	AccountIDFrom        string  `json:"account_id_from"`
	AccountIDTo          string  `json:"account_id_to"`
	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
}

type Message struct {
	Name    string                 `json:"Name"`
	Payload BalanceUpdatedInputDTO `json:"Payload"`
}

type ReplicateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
}

func NewReplicateAccountUseCase(accountGateway gateway.AccountGateway) *ReplicateAccountUseCase {
	return &ReplicateAccountUseCase{
		AccountGateway: accountGateway,
	}
}

func (uc *ReplicateAccountUseCase) Execute(input BalanceUpdatedInputDTO) error {
	accountFrom, err := entity.NewAccount(input.AccountIDFrom, input.BalanceAccountIDFrom)
	if err != nil {
		return err
	}
	err = SaveOrUpdate(uc.AccountGateway, *accountFrom)
	if err != nil {
		return err
	}

	accountTo, err := entity.NewAccount(input.AccountIDTo, input.BalanceAccountIDTo)
	if err != nil {
		return err
	}
	err = SaveOrUpdate(uc.AccountGateway, *accountTo)
	if err != nil {
		return err
	}

	return nil
}

func SaveOrUpdate(gateway gateway.AccountGateway, account entity.Account) error {
	_, err := gateway.FindByID(account.ID)
	if err != nil {
		err := gateway.Save(&account)
		if err != nil {
			return err
		}
		return nil
	}

	err = gateway.UpdateBalance(&account)
	if err != nil {
		return err
	}
	return nil
}
