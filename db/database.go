package db

import (
	"database/sql"
	"log"
)

// NewDBConnection returns a pointer to a SQL database
func NewDBConnection(databaseURL string) *sql.DB {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

// CreateTables creates tables in the database. This is an initial set up
func CreateTables(db *sql.DB) {
	CreateUserTable(db)
}
