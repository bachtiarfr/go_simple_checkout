package handlers

import "net/http"

type AuthenticationHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}

type UserHandler interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetUserByRefreshToken(w http.ResponseWriter, r *http.Request)
}
