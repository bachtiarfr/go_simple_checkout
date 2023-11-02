package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-checkout-app/internal/bootstrap"
	"simple-checkout-app/internal/config"
	handlers "simple-checkout-app/internal/handler"
	"simple-checkout-app/internal/repositories"
	"simple-checkout-app/internal/router"
	services "simple-checkout-app/internal/service"
)

func main() {
	config, err := config.ReadConfig("internal/config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return
	}

	db, errDb := bootstrap.InitDatabase(config)
	if errDb != nil {
		log.Fatalf("Failed to initialize the database: %v\n", errDb)
		return
	}

	userRepository := repositories.NewUserRepository(db)
	authService := services.NewAuthenticationService(userRepository)
	authHandler := handlers.NewAuthenticationHandler(authService)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	r := router.SetupRouter(authHandler, userHandler, productHandler)

	port := 8080
	fmt.Printf("Server is running on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
