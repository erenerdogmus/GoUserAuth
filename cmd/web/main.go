package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/erenerdogmus/internal/connection"
	_ "github.com/mattn/go-sqlite3"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Message  string `json:"message"`
}

func sendErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response := ErrorResponse{Code: code, Message: message}
	json.NewEncoder(w).Encode(response)
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./user.db")
	if err != nil {
		log.Fatalf("Veritabanı açılamadı: %v", err)
	}
	defer db.Close()
	
	connection.InitDb(db)

	log.Println("Sunucu başlatılıyor :8080")
	http.ListenAndServe(":8080", routes())
}
