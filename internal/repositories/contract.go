package repositories

import "simple-checkout-app/internal/entity"

type UserRepository interface {
	FindByEmail(email string) (*entity.User, error)
	FindAll() ([]entity.ListUser, error)
	CreateNewUser(user *entity.User) (entity.User, error)
	CreateOrUpdateRefreshToken(token string, userID int) error
	GetUserByRefreshToken(token string) (*entity.ListUser, error)
}

type ProductRepository interface {
	AddToCart(userID, productID int) error
	GetCart(userID int) ([]entity.Product, error)
	RemoveFromCart(userID, productID int) error
}
