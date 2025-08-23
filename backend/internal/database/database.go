package db

import (
	"os"
	"database/sql"
	"log"
)

func InitDb() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://postgres:password@localhost:5432/postgres?sslmode=disable"
	}

	var err error

	db,err := sql.Open("postgres", databaseURL)

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database: ", err)
	}

}

func createUsersTable() {
	query = `
	CREATE TABLE IF NOT EXISTS users {
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
		email VARCHAR(255) UNIQUE NULL,
		password_hash VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	}
	`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal("Failed to create users table: ", err)		
	}
}