package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simple-checkout-app/internal/entity"
	services "simple-checkout-app/internal/service"

	"golang.org/x/crypto/bcrypt"
)

type authenticationHandler struct {
	authService services.AuthenticationService
}

func NewAuthenticationHandler(authService services.AuthenticationService) AuthenticationHandler {
	return &authenticationHandler{authService}
}

func (h *authenticationHandler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		response entity.LoginResponse
		payload  struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
	)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	email := payload.Email
	password := payload.Password

	token, err := h.authService.Authenticate(email, password)
	if err != nil {
		response = entity.LoginResponse{
			Code:    "01",
			Message: fmt.Sprintf("failed : %s", err),
			Token:   "",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	response = entity.LoginResponse{
		Code:    "00",
		Message: "success",
		Token:   token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *authenticationHandler) Register(w http.ResponseWriter, r *http.Request) {
	var (
		response entity.RegisterResponse
		payload  struct {
			Fullname    string `json:"fullname"`
			Email       string `json:"email"`
			Age         int    `json:"age"`
			PhoneNumber string `json:"phone_number"`
			Password    string `json:"password"`
		}
	)

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	fullname := payload.Fullname
	email := payload.Email
	age := payload.Age
	phoneNumber := payload.PhoneNumber
	password := payload.Password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	user, err := h.authService.Register(fullname, email, age, phoneNumber, string(hashedPassword))
	if err != nil {
		response = entity.RegisterResponse{
			Code:    "01",
			Message: fmt.Sprintf("failed : %s", err),
			Data:    nil,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	response = entity.RegisterResponse{
		Code:    "00",
		Message: "success",
		Data:    &user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
