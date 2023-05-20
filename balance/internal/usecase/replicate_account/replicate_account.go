package replicateaccount

import (
	"context"

	"github.com/celsopires1999/fc-ms-balance/internal/entity"
	"github.com/celsopires1999/fc-ms-balance/internal/gateway"
	"github.com/celsopires1999/fc-ms-balance/pkg/uow"
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
	Uow uow.UowInterface
}

func NewReplicateAccountUseCase(uow uow.UowInterface) *ReplicateAccountUseCase {
	return &ReplicateAccountUseCase{
		Uow: uow,
	}
}

func (uc *ReplicateAccountUseCase) Execute(ctx context.Context, input BalanceUpdatedInputDTO) error {
	err := uc.Uow.Do(ctx, func(_ *uow.Uow) error {
		accountGateway := uc.getAccountGateway(ctx)

		accountFrom, err := entity.NewAccount(input.AccountIDFrom, input.BalanceAccountIDFrom)
		if err != nil {
			return err
		}
		err = uc.saveOrUpdate(accountGateway, *accountFrom)
		if err != nil {
			return err
		}

		accountTo, err := entity.NewAccount(input.AccountIDTo, input.BalanceAccountIDTo)
		if err != nil {
			return err
		}
		err = uc.saveOrUpdate(accountGateway, *accountTo)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (uc *ReplicateAccountUseCase) getAccountGateway(ctx context.Context) gateway.AccountGateway {
	repo, err := uc.Uow.GetRepository(ctx, "AccountDB")
	if err != nil {
		panic(err)
	}
	return repo.(gateway.AccountGateway)
}

func (uc *ReplicateAccountUseCase) saveOrUpdate(gateway gateway.AccountGateway, account entity.Account) error {
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
