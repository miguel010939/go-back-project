package db

import (
	"database/sql"
	"log"
)

const (
	createTableUsers = `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(30) NOT NULL,
    	hashedpswd CHAR(64) NOT NULL,
    	email VARCHAR(100) NOT NULL UNIQUE
	)`
	createTableSessions = `CREATE TABLE IF NOT EXISTS sessions (
		id SERIAL PRIMARY KEY,
		token VARCHAR(32) NOT NULL,
    	userx INTEGER NOT NULL,
    	createdat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	FOREIGN KEY (userx) REFERENCES users(id)
	)`
	createTableFollowers = `CREATE TABLE IF NOT EXISTS followers (
		id SERIAL PRIMARY KEY,
		usera INTEGER NOT NULL,
		userb INTEGER NOT NULL,
    	FOREIGN KEY (usera) REFERENCES users(id),
    	FOREIGN KEY (userb) REFERENCES users(id)
	)`
	createTableProducts = `CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		namex VARCHAR(50) NOT NULL,
    	description VARCHAR(300),
    	imageurl VARCHAR(150) NOT NULL,
    	userx INTEGER NOT NULL,
    	FOREIGN KEY (userx) REFERENCES users(id)
	)`
	createTableFavorites = `CREATE TABLE IF NOT EXISTS favorites (
		id SERIAL PRIMARY KEY,
		userx INTEGER NOT NULL,
		product INTEGER NOT NULL,
		FOREIGN KEY (userx) REFERENCES users(id),
    	FOREIGN KEY (product) REFERENCES products(id)
	)`
	createTableBids = `CREATE TABLE IF NOT EXISTS bids (
		id SERIAL PRIMARY KEY,
		userx INTEGER NOT NULL,
		product INTEGER NOT NULL,
		amount DECIMAL(8,2) NOT NULL,
    	createdat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	FOREIGN KEY (userx) REFERENCES users(id),
    	FOREIGN KEY (product) REFERENCES products(id)
	)`
)

const (
	deleteTableUsers     = "DROP TABLE IF EXISTS users CASCADE"
	deleteTableSessions  = "DROP TABLE IF EXISTS sessions CASCADE"
	deleteTableFollowers = "DROP TABLE IF EXISTS followers CASCADE"
	deleteTableProducts  = "DROP TABLE IF EXISTS products CASCADE"
	deleteTableFavorites = "DROP TABLE IF EXISTS favorites CASCADE"
	deleteTableBids      = "DROP TABLE IF EXISTS bids CASCADE"
)

func execQuery(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
