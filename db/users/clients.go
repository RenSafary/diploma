package db

import (
	"database/sql"
	"diploma/utils"
	"log"
	"strconv"
)

type Clients struct {
	DB *sql.DB
}

func ClientsInit(db *sql.DB) *Clients {
	return &Clients{DB: db}
}

type Client struct {
	Id       int
	Username string
	Password string
}

func (d *Clients) GetClient(username, password string) (string, error) {
	client := &Client{}

	err := d.DB.QueryRow(
		"SELECT id, passwd FROM clients WHERE username=$1", username,
	).Scan(&client.Id, &client.Password)

	if err == sql.ErrNoRows {
		log.Println("Client does not exist: ", err)
		return "", err
	} else if err != nil {
		log.Println(err)
		return "", err
	}

	// comparing password
	err = utils.CompareHashPass(client.Password, password)
	if err != nil {
		return "Wrong login or password", nil
	}

	// giving jwt token
	token, err := utils.GenerateToken(client.Id, username)
	if err != nil {
		return "Wrong login or password", nil
	}

	return token, nil
}

func (c *Clients) CreateClient(username, password, firstname, lastname, email, sex, age string) (string, error) {
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
		`INSERT INTO clients (Username, Passwd, FirstName, LastName, Email, Sex, Age)
         VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		username, hashedPass, firstname, lastname, email, sex, age_int,
	)
	if err != nil {
		log.Println("Error inserting client:", err)
		return "", err
	}

	return "Hello, world!", nil
}
