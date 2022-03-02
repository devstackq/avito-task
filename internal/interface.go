package internal

import "avito/internal/model"

type AccountBalanceRepositoryInterface interface {
	Add(account *model.Account) error
	Debit(account *model.Account) error
	CreateAccount(uuid string) (int64, error)
	CheckBalance(uuid string, currencyType int) (float64, error)
}

type AccountBalanceServiceInterface interface {
	Add(account *model.Account) error
	Debit(account *model.Account) error
	CheckBalance(uuid string, currencyType int) (float64, error)
	NewAccount() (int64, error)

	Transfer(sender model.Account, receiver model.Account) error
}

type UserServiceInterface interface {
	IsExistUser(id int) (string, error)
	CreateUser(user *model.User) (int64, error)
}
type UserRepoInterface interface {
	IsExistUser(id int) (string, error)
	CreateUser(user *model.User) (int64, error)
}

type CurrencyServiceInterface interface {
	Create(string) error
}
type CurrencyRepositoryInterface interface {
	GetCurrencyID(string) (int, error)
	Create(string) error
}
