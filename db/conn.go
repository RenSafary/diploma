package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type ClinicDB struct {
	DB *sql.DB
}

func getEnvVariablesDB() string {
	if err := godotenv.Load(); err != nil {
		log.Println("Couldn't get env variables for DB", err)
		return ""
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", user, password, db_name, host, port, sslmode)

	return connStr
}

func Conn() (*ClinicDB, error) {
	connStr := getEnvVariablesDB()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Couldn't connect to DB", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Println("DB does not work", err)
		return nil, err
	}

	return &ClinicDB{DB: db}, nil
}

func (d *ClinicDB) GetClient(username, password string) (string, error) {
	var db_pass string
	err := d.DB.QueryRow(
		"SELECT password FROM users WHERE username=$1", username,
	).Scan(&db_pass)

	if err == sql.ErrNoRows {
		log.Println("User does not exist: ", err)
		return "", err
	} else if err != nil {
		log.Println("Wrong user's password: ", err)
		return "", err
	}

	return "Hello, world!", nil // instead of the token before development
}
