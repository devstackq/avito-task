package repository

import (
	"avito/internal"
	"avito/internal/model"
	"database/sql"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) internal.UserRepoInterface {
	return &UserRepo{db: db}
}
func (ur *UserRepo) IsExistUser(id int) (uuid string, err error) {

	query := `SELECT uuid FROM users left join account a on a.id = users.account_id where user_id=$1`

	err = ur.db.QueryRow(query, id).Scan(&uuid)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

func (ur *UserRepo) CreateUser(user *model.User) (res int64, err error) {

	query := `INSERT INTO users(email, name, account_id, password) VALUES ($1, $2, $3, $4) RETURNING user_id`
	err = ur.db.QueryRow(query,
		user.Email,
		user.Name,
		user.Account.ID,
		user.Password,
	).Scan(&res)

	if err != nil {
		return 0, err
	}
	return res, nil
}
