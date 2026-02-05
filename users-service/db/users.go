package db

import (
	"database/sql"
	"diploma/auth-service/utils"
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
	Id        int
	Username  string
	Password  string
	Email     string
	FirstName string
	LastName  string
	Sex       string
	Admin     bool
}

func (c *Users) CreateUser(username, password, firstname, lastname, email, sex, age string) (int32, error) {
	age_int, err := strconv.Atoi(age)
	if err != nil {
		return 0, err
	}

	// making hashed password
	hashedPass, err := utils.MakeHashed(password)
	if err != nil {
		return 0, err
	}

	var userId int32
	err = c.DB.QueryRow(
		`INSERT INTO users (Username, Passwd, FirstName, LastName, Email, Sex, Age, Adm)
         VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id`,
		username, hashedPass, firstname, lastname, email, sex, age_int, false,
	).Scan(&userId)
	if err != nil {
		log.Println("Error inserting user:", err)
		return 0, err
	}

	return userId, nil
}
