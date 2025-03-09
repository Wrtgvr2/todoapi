package repository

import (
	"database/sql"
	"fmt"
	"os"

	env "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func OpenDatabase() (*sql.DB, error) {
	if err := env.Load(); err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}
	if err = DB.Ping(); err != nil {
		return nil, err
	}

	return DB, nil
}

func CloseDatabase() error {
	return DB.Close()
}
