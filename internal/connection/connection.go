package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func InitDb(db *sql.DB) {
	if err := migrate(db); err != nil {
		fmt.Printf("Failed to migrate database: %v\n", err)
	}
}

func migrate(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}
