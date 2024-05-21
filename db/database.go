package db

import (
	"database/sql"
)

// NewDBConnection returns a pointer to a SQL database
func NewDBConnection(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// CreateTables creates tables in the database. This is an initial set up
func CreateTables(db *sql.DB) {
	execQuery(db, createTableUsers)
	execQuery(db, createTableSessions)
	execQuery(db, createTableFollowers)
	execQuery(db, createTableProducts)
	execQuery(db, createTableFavorites)
	execQuery(db, createTableBids)
}

func DropTables(db *sql.DB) {
	execQuery(db, deleteTableUsers)
	execQuery(db, deleteTableSessions)
	execQuery(db, deleteTableFollowers)
	execQuery(db, deleteTableProducts)
	execQuery(db, deleteTableFavorites)
	execQuery(db, deleteTableBids)
}
