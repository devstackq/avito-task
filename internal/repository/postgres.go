package repository

import (
	"avito/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/stdlib"
	// _ "github.com/jackc/pgx/v4"
)

func NewPostgres(cfg *config.DBConf) (*sql.DB, error) {
	dbURI := fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Port,
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.DBName,
	)

	db, err := sql.Open(cfg.Dialect, dbURI)
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		log.Println("couldn't ping: postgres", err)
		return nil, err
	}

	return db, nil
}

func CreateTables(db *sql.DB) error {

	currencies, err := db.Prepare(`CREATE TABLE  IF NOT EXISTS currency(
		currency_id SERIAL PRIMARY KEY,
		name VARCHAR(20) NOT NULL,
		UNIQUE(name)
	);`)
	if err != nil {
		return err
	}
	_, err = currencies.Exec()
	if err != nil {
		return err
	}
	accounts, err := db.Prepare(`CREATE TABLE IF NOT EXISTS account(
		account_id SERIAL PRIMARY KEY,
		uuid VARCHAR(200) NOT NULL,
		balance DECIMAL,
		currency_type INTEGER REFERENCES currency(currency_id)
	);`)
	if err != nil {
		return err
	}
	_, err = accounts.Exec()
	if err != nil {
		return err
	}
	users, err := db.Prepare(`CREATE TABLE IF NOT EXISTS users(
		user_id SERIAL PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		email VARCHAR(70) NOT NULL,
		password VARCHAR(50) NOT NULL,
		account_id INTEGER REFERENCES account(account_id),
		UNIQUE(email)
	);`)

	if err != nil {
		return err
	}
	_, err = users.Exec()
	if err != nil {
		return err
	}
	log.Println("success create tables")
	return nil
}

// func AddCurrency(name string, db *sql.DB) (err error) {
// 	query := `INSERT INTO currency(name) VALUES ($1)`
// 	_, err = db.Exec(query, name)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
