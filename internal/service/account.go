package service

import (
	"avito/internal"
	"avito/internal/model"

	"github.com/gofrs/uuid"
)

type AccountService struct {
	accountRepo internal.AccountBalanceRepositoryInterface
}

func NewAccountService(repo internal.AccountBalanceRepositoryInterface) internal.AccountBalanceServiceInterface {
	return &AccountService{accountRepo: repo}
}

func (a *AccountService) Add(account *model.Account) error {
	return a.accountRepo.Add(account)
}

func (a *AccountService) Debit() error {
	return nil
}

func (a *AccountService) CheckBalanceByID() (bool, error) {
	return true, nil

}

func (a *AccountService) UpdateBalanceByID() error {
	return nil

}
func (a *AccountService) NewAccount() (int64, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return 0, err
	}
	// user.UUID = id.String()
	return a.accountRepo.CreateAccount(id.String())
}
