package repository

import (
	"avito/internal"
	"avito/internal/model"
	"database/sql"
	"log"
)

type AccountRepo struct {
	db *sql.DB
}

func NewAccountRepo(db *sql.DB) internal.AccountBalanceRepositoryInterface {
	return &AccountRepo{db: db}
}

func (ar *AccountRepo) CreateAccount(uuid string) (res int64, err error) {

	query := `INSERT INTO account(uuid, balance, currency_type) VALUES ($1, $2, $3) RETURNING account_id`

	err = ar.db.QueryRow(query, uuid, 0, 1).Scan(&res)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (ar *AccountRepo) Add(account *model.Account) error {
	log.Print(account, "add")
	query := "UPDATE account SET balance = balance + $1 WHERE uuid = $2 AND currency_type=$3"
	_, err := ar.db.Exec(query, account.WalletAmount, account.UUID, account.CurrencyType)
	if err != nil {
		return nil
	}
	return nil
}

//in further..
//AddNewCurrency(uuid)

func (ar *AccountRepo) CheckBalanceByUUID(uuid string) (amount float64, err error) {
	query := `SELECT balance FROM account WHERE uuid = $1`
	err = ar.db.QueryRow(query, uuid).Scan(&amount)
	if err != nil {
		return 0, err
	}
	return

}
func (ar *AccountRepo) Debit(account *model.Account) error {
	query := "UPDATE account SET balance = balance - $1 WHERE uuid = $2 AND currency_type=$3"
	_, err := ar.db.Exec(query, account.WalletAmount, account.UUID, account.CurrencyType)
	if err != nil {
		return nil
	}
	return nil
}

// func (ar *AccountRepo) Transfer() error {
// 	return nil
// }
