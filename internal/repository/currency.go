package repository

import (
	"avito/internal"
	"database/sql"
)

type CurrencyRepo struct {
	db *sql.DB
}

func NewCurrencyRepo(db *sql.DB) internal.CurrencyRepositoryInterface {
	return &CurrencyRepo{db: db}
}

func (cr CurrencyRepo) GetCurrencyID(name string) (res int, err error) {
	query := `SELECT currency_id FROM currency WHERE name = $1`
	err = cr.db.QueryRow(query, name).Scan(&res)
	if err != nil {
		return 0, err
	}
	return
}
