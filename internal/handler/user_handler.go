package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"simple-checkout-app/internal/entity"
	services "simple-checkout-app/internal/service"
)

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var response entity.UserResponse

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		log.Printf("Token missing")
		response = entity.UserResponse{
			Code:    "01",
			Message: "faield : token missing",
			Data:    nil,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	_, err := h.userService.ValidateToken(tokenString)
	if err != nil {
		log.Printf("Token not validated :", err)
		response = entity.UserResponse{
			Code:    "01",
			Message: fmt.Sprintf("token validation error : %s", err),
			Data:    nil,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	users, err := h.userService.GetAllUsers()
	if err != nil {
		log.Printf("Failed to get user data :", err)
		response = entity.UserResponse{
			Code:    "00",
			Message: fmt.Sprintf("failed : %s", err),
			Data:    users,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	response = entity.UserResponse{
		Code:    "00",
		Message: "success",
		Data:    users,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *userHandler) GetUserByRefreshToken(w http.ResponseWriter, r *http.Request) {

	var response entity.UserAffiliatedResponse

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		log.Printf("Token missing")
		response = entity.UserAffiliatedResponse{
			Code:    "01",
			Message: "faield : token missing",
			Data:    nil,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := h.userService.GetUserByRefreshToken(tokenString)
	if err != nil {
		log.Printf("Failed to get user data :", err)
		response = entity.UserAffiliatedResponse{
			Code:    "01",
			Message: fmt.Sprintf("token validation error : %s", err),
			Data:    nil,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	response = entity.UserAffiliatedResponse{
		Code:    "00",
		Message: "success",
		Data:    user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
