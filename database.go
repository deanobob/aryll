package main

import (
	"database/sql"

	// blank import due to go-sqlite3 using standard golang database/sql implementation
	_ "github.com/mattn/go-sqlite3"
)

// Connect creates a connection to the given sqlite3 file
func Connect(database string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		return db, db.Ping()
	}

	return db, err
}
