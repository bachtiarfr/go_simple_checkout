package services

import (
	"simple-checkout-app/internal/entity"

	"github.com/dgrijalva/jwt-go"
)

type AuthenticationService interface {
	Authenticate(email, password string) (string, error)
	Register(fullname string, email string, age int, phoneNumber string, hashedPassword string) (entity.User, error)
}

type UserService interface {
	GetAllUsers() ([]entity.ListUser, error)
	ValidateToken(tokenString string) (jwt.Claims, error)
	GetUserByRefreshToken(token string) (*entity.ListUser, error)
}

type ProductService interface {
	AddToCart(userID, productID int) error
	GetCart(userID int) ([]entity.Product, error)
	RemoveFromCart(userID, productID int) error
}
