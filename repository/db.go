package repository

import (
	"database/sql"
	"fmt"
	"os"

	env "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func OpenDatabase() error {
	if err := env.Load(); err != nil {
		return err
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
		return err
	}

	if err := DB.Ping(); err != nil {
		return err
	}

	if err := tablesInit(); err != nil {
		return err
	}

	return nil
}

func CloseDatabase() {
	DB.Close()
}

func tablesInit() error {
	createTables := `
	CREATE TABLE IF NOT EXISTS users (
		id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		username VARCHAR(24) NOT NULL,
		password VARCHAR(60) NOT NULL
	);
	CREATE TABLE IF NOT EXISTS todos (
		id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		user_id BIGINT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
		title VARCHAR(100) NOT NULL,
		description TEXT,
		completed BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := DB.Exec(createTables); err != nil {
		return err
	}

	return nil
}
