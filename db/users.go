package db

import (
	"database/sql"
	"diploma/utils"
	"log"
	"strconv"
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
}

func (d *Users) GetUser(username, password string) (string, error) {
	user := &User{}

	err := d.DB.QueryRow(
		"SELECT id, passwd FROM users WHERE username=$1", username,
	).Scan(&user.Id, &user.Password)

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
	token, err := utils.GenerateToken(user.Id, username)
	if err != nil {
		return "Wrong login or password", nil
	}

	return token, nil
}

func (c *Users) CreateUser(username, password, firstname, lastname, email, sex, age string) (string, error) {
	age_int, err := strconv.Atoi(age)
	if err != nil {
		return "", err
	}

	// making hashed password
	hashedPass, err := utils.MakeHashed(password)
	if err != nil {
		return "", err
	}

	_, err = c.DB.Exec(
		`INSERT INTO users (Username, Passwd, FirstName, LastName, Email, Sex, Age, Adm)
         VALUES ($1,$2,$3,$4,$5,$6,$7, $8)`,
		username, hashedPass, firstname, lastname, email, sex, age_int, false,
	)
	if err != nil {
		log.Println("Error inserting user:", err)
		return "", err
	}

	return "Hello, world!", nil
}
