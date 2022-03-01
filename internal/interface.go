package internal

import "avito/internal/model"

type AccountBalanceRepositoryInterface interface {
	Add(account *model.Account) error
	Debit() error
	CreateAccount(uuid string) (int64, error)
}

type AccountBalanceServiceInterface interface {
	Add(account *model.Account) error
	Debit() error
	CheckBalanceByID() (bool, error)
	UpdateBalanceByID() error
	NewAccount() (int64, error)
}

type UserServiceInterface interface {
	IsExistUser(id int) (string, error)
	CreateUser(user *model.User) (int64, error)
}
type UserRepoInterface interface {
	IsExistUser(id int) (string, error)
	CreateUser(user *model.User) (int64, error)
}
