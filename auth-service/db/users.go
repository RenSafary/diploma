package db

import (
	"database/sql"
	"diploma/auth-service/utils"
	"log"
)

type Users struct {
	DB *sql.DB
}

func UsersInit(db *sql.DB) *Users {
	return &Users{DB: db}
}

type User struct {
	Id       int
	Username string
	Password string
	Admin    bool
}

func (d *Users) GiveToken(username, password string) (string, error) {
	user := &User{}

	err := d.DB.QueryRow(
		"SELECT id, passwd, adm FROM users WHERE username=$1", username,
	).Scan(&user.Id, &user.Password, &user.Admin)

	if err == sql.ErrNoRows {
		log.Println("User does not exist: ", err)
		return "", err
	} else if err != nil {
		log.Println(err)
		return "", err
	}

	// comparing password
	err = utils.CompareHashPass(user.Password, password)
	if err != nil {
		return "Wrong login or password", nil
	}

	// giving jwt token
	token, err := utils.GenerateToken(user.Id, username, user.Admin)
	if err != nil {
		return "Wrong login or password", nil
	}

	return token, nil
}
