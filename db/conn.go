package db

import (
	"database/sql"
	db_cli "diploma/db/users"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type ClinicDB struct {
	DB      *sql.DB
	Clients *db_cli.Clients // struct to control the table
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

	createTable := `
	CREATE TABLE IF NOT EXISTS clients (
		Id SERIAL PRIMARY KEY,
		Username VARCHAR(250),
		Passwd VARCHAR(250),
		FirstName VARCHAR(100),
		LastName VARCHAR(100),
		Email VARCHAR(250),
		Age INTEGER,
		Sex CHAR(1)
	);
	`
	if _, err := db.Exec(createTable); err != nil {
		log.Println("Couldn't create clients table:", err)
		return nil, err
	}

	return &ClinicDB{
		DB:      db,
		Clients: db_cli.ClientsInit(db),
	}, nil
}
