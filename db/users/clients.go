package db

import (
	"database/sql"
	"log"
	"strconv"
)

type Clients struct {
	DB *sql.DB
}

func ClientsInit(db *sql.DB) *Clients {
	return &Clients{DB: db}
}

func (d *Clients) GetClient(username, password string) (string, error) {
	var db_pass string
	err := d.DB.QueryRow(
		"SELECT passwd FROM clients WHERE username=$1", username,
	).Scan(&db_pass)

	if err == sql.ErrNoRows {
		log.Println("Client does not exist: ", err)
		return "", err
	} else if err != nil {
		log.Println(err)
		return "", err
	}

	if db_pass != password {
		return "Wrong login or password", nil
	}

	return "Hello, world!", nil // instead of the token before development
}

func (c *Clients) CreateClient(username, password, firstname, lastname, email, sex, age string) (string, error) {
	age_int, err := strconv.Atoi(age)
	if err != nil {
		return "", err
	}

	_, err = c.DB.Exec(
		`INSERT INTO clients (Username, Passwd, FirstName, LastName, Email, Sex, Age)
         VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		username, password, firstname, lastname, email, sex, age_int,
	)
	if err != nil {
		log.Println("Error inserting client:", err)
		return "", err
	}
	return "Hello, world!", nil
}
