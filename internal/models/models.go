package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/erenerdogmus/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

func InsertUser(db *sql.DB, username, email, hashedPassword string) (int64, error) {
	statement := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`
	result, err := db.Exec(statement, username, email, hashedPassword)
	if err != nil {
		fmt.Printf("Error inserting user: %v\n", err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func Authenticate(db *sql.DB, email, password string) (int, string, error) {
	var id int
	var username, hashedPassword string
	statement := `SELECT id, username, password FROM users WHERE email=?`
	err := db.QueryRow(statement, email).Scan(&id, &username, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, "", validator.ErrInvalidCredentials
		}
		return 0, "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, "", validator.ErrInvalidCredentials
		}
		return 0, "", err
	}
	return id, username, nil
}
