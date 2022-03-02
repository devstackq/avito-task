package service

import (
	"avito/internal"
	"avito/internal/model"
	"fmt"

	"github.com/gofrs/uuid"
)

type AccountService struct {
	accountRepo  internal.AccountBalanceRepositoryInterface
	currencyRepo internal.CurrencyRepositoryInterface
}

func NewAccountService(accountRepo internal.AccountBalanceRepositoryInterface, currencyRepo internal.CurrencyRepositoryInterface) internal.AccountBalanceServiceInterface {
	return &AccountService{accountRepo: accountRepo,
		currencyRepo: currencyRepo}
}

func (as *AccountService) CheckBalance(uuid string, currencyType int) (float64, error) {
	return as.accountRepo.CheckBalance(uuid, currencyType)
}

func (as *AccountService) Add(account *model.Account) (err error) {
	account.CurrencyType, err = as.currencyRepo.GetCurrencyID(account.Currency)
	if err != nil {
		return err
	}
	return as.accountRepo.Add(account)
}

func (as *AccountService) Debit(account *model.Account) (err error) {
	account.CurrencyType, err = as.currencyRepo.GetCurrencyID(account.Currency)
	if err != nil {
		return err
	}
	balanceAmount, err := as.CheckBalance(account.UUID, account.CurrencyType)
	if err != nil {
		return err
	}
	if balanceAmount-account.WalletAmount < 1 {
		return fmt.Errorf("nedostatochno sredstv")
	}
	return as.accountRepo.Debit(account)
}

func (as *AccountService) NewAccount() (int64, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return 0, err
	}
	return as.accountRepo.CreateAccount(id.String())
}

func (as *AccountService) Transfer(sender model.Account, receiver model.Account) (err error) {
	if err = as.Debit(&sender); err != nil {
		return err
	}
	if err = as.Add(&receiver); err != nil {
		return err
	}
	return nil
}
