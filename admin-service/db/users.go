package db

import (
	"database/sql"
)

type Users struct {
	DB *sql.DB
}

func UsersInit(db *sql.DB) *Users {
	return &Users{DB: db}
}

func (u *Users) MakeAdminDB(user_id int) error {
	_, err := u.DB.Exec(
		"UPDATE users SET Adm = $1 WHERE Id = $2",
		true,
		user_id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *Users) DeleteAdminDB(user_id int) error {
	_, err := u.DB.Exec(
		"UPDATE users SET Adm = $1 WHERE Id = $2",
		false,
		user_id,
	)
	if err != nil {
		return err
	}
	return nil
}
