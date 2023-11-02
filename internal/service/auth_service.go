package services

import (
	"fmt"
	"log"
	"simple-checkout-app/internal/entity"
	"simple-checkout-app/internal/repositories"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type authenticationService struct {
	userRepository repositories.UserRepository
}

func NewAuthenticationService(userRepository repositories.UserRepository) AuthenticationService {
	return &authenticationService{
		userRepository: userRepository,
	}
}

func (s *authenticationService) Authenticate(email, password string) (string, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if user == nil || !s.verifyPassword(user, password) {
		log.Printf("invalid email or password: %v %v", email, password)
		return "", fmt.Errorf("invalid email or password")
	}

	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Minute * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("my-secret-key"))

	errToken := s.userRepository.CreateOrUpdateRefreshToken(tokenString, user.ID)
	if errToken != nil {
		return "", errToken
	}

	return tokenString, err
}

func (s *authenticationService) Register(fullname string, email string, age int, phoneNumber string, hashedPassword string) (entity.User, error) {
	user, err := s.userRepository.FindByEmail(email)
	if user != nil {
		log.Printf("Email exist: %v", err)
		return entity.User{}, fmt.Errorf("Email exist")
	}

	dataUser := entity.User{
		Fullname:    fullname,
		Email:       email,
		Age:         age,
		PhoneNumber: phoneNumber,
		Password:    hashedPassword,
	}

	newUser, err := s.userRepository.CreateNewUser(&dataUser)
	if user != nil {
		log.Printf("Failed when create user: %v", err)
		return entity.User{}, fmt.Errorf("Email exist")
	}
	fmt.Printf("newUser :", newUser)
	return entity.User{}, err
}

func (s *authenticationService) verifyPassword(user *entity.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
