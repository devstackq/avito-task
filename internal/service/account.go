package service

import (
	"avito/internal"
	"avito/internal/model"
	"fmt"
	"log"

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

func (as *AccountService) CheckBalanceByUUID(uuid string) (float64, error) {
	return as.accountRepo.CheckBalanceByUUID(uuid)
}

func (as *AccountService) Add(account *model.Account) error {
	var err error
	account.CurrencyType, err = as.currencyRepo.GetCurrencyID(account.Currency)
	log.Println(account, err)

	return as.accountRepo.Add(account)
}

func (as *AccountService) Debit(account *model.Account) error {
	account.CurrencyType, _ = as.currencyRepo.GetCurrencyID(account.Currency)

	return as.accountRepo.Debit(account)
}

func (as *AccountService) NewAccount() (int64, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return 0, err
	}
	return as.accountRepo.CreateAccount(id.String())
}

func (as *AccountService) Transfer(sender model.Account, receiver model.Account) error {

	var err error

	sender.WalletAmount, err = as.accountRepo.CheckBalanceByUUID(sender.UUID)
	if err != nil {
		return err
	}

	//70; 70 - 10 ?
	if sender.WalletAmount-sender.TransferAmount > 0 {
		// GetCurrencyID(balanceSender.Currency)
		sender.CurrencyType, err = as.currencyRepo.GetCurrencyID(sender.Currency)

		sender.WalletAmount = sender.TransferAmount // 10
		if err = as.Debit(&sender); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("ne dostatochno deneg")
	}

	receiver.CurrencyType = sender.CurrencyType

	if err = as.Add(&receiver); err != nil {
		return err
	}
	return nil
}
