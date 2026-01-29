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
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &ClinicDB{DB: db}, nil
}
