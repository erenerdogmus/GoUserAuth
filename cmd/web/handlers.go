package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/erenerdogmus/internal/models"
	"github.com/erenerdogmus/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirm := r.FormValue("confirm")

	if !validator.NotBlank(username) {
		sendErrorResponse(w, http.StatusBadRequest, "Kullanıcı adı boş olamaz")
		return
	} else if !validator.NotBlank(email) {
		sendErrorResponse(w, http.StatusBadRequest, "Email boş olamaz")
		return
	} else if !validator.NotBlank(password) {
		sendErrorResponse(w, http.StatusBadRequest, "Şifre boş olamaz")
		return
	} else if !validator.NotBlank(confirm) {
		sendErrorResponse(w, http.StatusBadRequest, "Şifre boş olamaz")
		return
	}

	if !validator.Matches(username, validator.UsernameRX) {
		sendErrorResponse(w, http.StatusBadRequest, "Geçerli bir kullanıcı adı girin")
		return
	} else if !validator.Matches(email, validator.EmailRX) {
		sendErrorResponse(w, http.StatusBadRequest, "Geçerli bir email adresi girin")
		return
	}

	if !validator.IsUsernameUnique(db, username) {
		sendErrorResponse(w, http.StatusBadRequest, "Benzersiz bir kullanıcı adı girin")
		return
	} else if !validator.IsEmailUnique(db, email) {
		sendErrorResponse(w, http.StatusBadRequest, "Benzersiz bir email adresi girin.")
		return
	}

	if password != confirm {
		sendErrorResponse(w, http.StatusBadRequest, "Şifreler uyuşmuyor")
		return
	}

	if !validator.MinChars(password, 8) {
		sendErrorResponse(w, http.StatusBadRequest, "Şifre en az 8 karakter olmalıdır")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Şifre oluşturulurken hata oluştu")
		return
	}

	id, err := models.InsertUser(db, username, email, string(hashedPassword))
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Kullanıcı kaydı sırasında hata oluştu")
		return
	}

	user := User{
		ID:       id,
		Username: username,
		Email:    email,
		Message: "Kayıt işlemi Başarılı",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")

	if !validator.NotBlank(email) {
		sendErrorResponse(w, http.StatusBadRequest, "Email boş olamaz")
		return
	} else if !validator.NotBlank(password) {
		sendErrorResponse(w, http.StatusBadRequest, "Şifre boş olamaz")
		return
	}

	id, username, err := models.Authenticate(db, email, password)
	if err != nil {
		if errors.Is(err, validator.ErrInvalidCredentials) {
			sendErrorResponse(w, http.StatusUnauthorized, "Geçersiz e-posta veya şifre")
		} else {
			sendErrorResponse(w, http.StatusInternalServerError, "Giriş sırasında bir hata oluştu")
		}
		return
	}

	user := User{
		ID:       int64(id),
		Username: username,
		Email:    email,
		Message:  "Giriş başarılı",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
