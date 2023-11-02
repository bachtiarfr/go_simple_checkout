package services

import (
	"log"
	"simple-checkout-app/internal/bootstrap"
	"simple-checkout-app/internal/config"
	"simple-checkout-app/internal/repositories"
	services "simple-checkout-app/internal/service"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthService_Authenticate(t *testing.T) {
	config, err := config.ReadConfig("../config/config.yaml")
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

	token, err := authService.Authenticate("test@example.com", "password")

	assert.NoError(t, err)
	assert.NotNil(t, token)
}

func TestUserService_GetAllUsers(t *testing.T) {
	config, err := config.ReadConfig("../config/config.yaml")
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
	userService := services.NewUserService(userRepository)

	users, err := userService.GetAllUsers()

	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, 2, len(users))
}
