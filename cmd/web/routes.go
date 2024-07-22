package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/user/signup", Signup)
	router.HandlerFunc(http.MethodPost, "/user/login", Login)
	return router
}
