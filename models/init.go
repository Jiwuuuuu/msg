package models

import (
	"database/sql"
	"fmt"
)

// InitDB is used to initialize the database
func InitDB(driver, path string) (*sql.DB, error) {
	db, err := sql.Open(driver, path)
	if err != nil {
		return nil, fmt.Errorf("failed to open a connection to the database: %w", err)
	}

	statements := []string{
		`
		CREATE TABLE IF NOT EXISTS messages (
			messageID INTEGER PRIMARY KEY AUTOINCREMENT,
			roomID INTEGER,
			timestamp INTEGER,
			content TEXT NOT NULL,
			checksum TEXT NOT NULL,
			username TEXT NOT NULL,
		);
		`,
		`
		CREATE TABLE IF NOT EXISTS rooms (
			roomID INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
		);
		`,
	}

	for _, s := range statements {
		if _, err := db.Exec(s); err != nil {
			return nil, fmt.Errorf("failed to create the table: %w", err)
		}
	}

	return db, nil
}
