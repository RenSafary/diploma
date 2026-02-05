package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type ClinicDB struct {
	DB    *sql.DB
	Users *Users
}

func getEnvVariablesDB() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s search_path=public", user, password, db_name, host, port, sslmode)

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

	return &ClinicDB{
		DB:    db,
		Users: UsersInit(db),
	}, nil
}
