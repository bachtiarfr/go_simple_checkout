package services

import (
	"fmt"
	"log"
	"simple-checkout-app/internal/config"
	"simple-checkout-app/internal/entity"
	"simple-checkout-app/internal/repositories"

	"github.com/dgrijalva/jwt-go"
)

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) ValidateToken(tokenString string) (jwt.Claims, error) {
	config, err := config.ReadConfig("internal/config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func (s *userService) GetAllUsers() ([]entity.ListUser, error) {
	return s.userRepository.FindAll()
}

func (s *userService) GetUserByRefreshToken(token string) (*entity.ListUser, error) {
	return s.userRepository.GetUserByRefreshToken(token)
}
