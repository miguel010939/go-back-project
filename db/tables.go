package db

import (
	"database/sql"
	"log"
)

func CreateUserTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(30) NOT NULL
	)`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
