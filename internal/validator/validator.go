package validator

import (
	"database/sql"
	"errors"
	"regexp"
	"strings"
	"unicode/utf8"
)

var ErrInvalidCredentials = errors.New("geçersiz e-posta veya şifre")
var EmailRX = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
var UsernameRX = regexp.MustCompile(`^[a-zA-Z0-9._%+-]{3,20}$`)

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(strings.TrimSpace(value))
}

func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

func IsUsernameUnique(db *sql.DB, username string) bool {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE username = ?)"
	err := db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false
	}
	return !exists
}

func IsEmailUnique(db *sql.DB, email string) bool {
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE email = ?)"
	err := db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false
	}
	return !exists
}
