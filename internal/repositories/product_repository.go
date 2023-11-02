package repositories

import (
	"database/sql"
	"simple-checkout-app/internal/entity"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) AddToCart(userID, productID int) error {
	return nil
}

func (r *productRepository) GetCart(userID int) ([]entity.Product, error) {
	var products []entity.Product
	return products, nil
}

func (r *productRepository) RemoveFromCart(userID, productID int) error {
	return nil
}
