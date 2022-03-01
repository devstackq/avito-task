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
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (ar *AccountRepo) Add(account *model.Account) error {
	// select(balance where uuid=$1)
	log.Println(account, 123)

	query := "UPDATE account SET balance = balance + $1 WHERE id = $2"
	_, err := ar.db.Exec(query, account.WalletAmount, account.UUID)
	if err != nil {
		return nil
	}
	return nil
}

//in further..
// func (ar *AccountRepo) GetBalanceByUUID(uuid string) (int, error) {
// 	query := `SELECT walletAmount FROM account where uuid = $1`
// }

func (a *AccountRepo) Debit() error {
	return nil
}

func (a *AccountRepo) Transfer() error {
	return nil
}
