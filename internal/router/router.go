package router

import (
	handlers "simple-checkout-app/internal/handler"

	"github.com/gorilla/mux"
)

func SetupRouter(
	authHandler handlers.AuthenticationHandler,
	userHandler handlers.UserHandler,
	productHandler handlers.ProductHandler,
) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/api/v1/register", authHandler.Register).Methods("POST")

	// user endpoint
	r.HandleFunc("/api/v1/users", userHandler.GetAllUsers).Methods("GET")
	r.HandleFunc("/api/v1/user", userHandler.GetUserByRefreshToken).Methods("GET")

	// product endpoint
	r.HandleFunc("/api/v1/add-to-cart", productHandler.AddToCart).Methods("POST")

	return r
}
