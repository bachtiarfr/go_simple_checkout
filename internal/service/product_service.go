package services

import (
	"simple-checkout-app/internal/entity"
	"simple-checkout-app/internal/repositories"
)

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) ProductService {
	return &productService{productRepository}
}

func (s *productService) AddToCart(userID, productID int) error {
	// Implementasi logika untuk menambahkan produk ke keranjang
	// Anda dapat memeriksa apakah pengguna memiliki keranjang, jika belum, buat keranjang baru
	// Selanjutnya, tambahkan produk ke dalam keranjang
	// Misalnya, Anda dapat menggunakan metode dari productRepository
	return nil
}

func (s *productService) GetCart(userID int) ([]entity.Product, error) {
	// Implementasi logika untuk mendapatkan daftar produk dalam keranjang
	// Anda perlu mengambil daftar produk yang ada di keranjang berdasarkan userID
	// Misalnya, Anda dapat menggunakan metode dari productRepository
	return nil, nil
}

func (s *productService) RemoveFromCart(userID, productID int) error {
	// Implementasi logika untuk menghapus produk dari keranjang
	// Anda perlu memeriksa apakah produk ada di keranjang pengguna dan menghapusnya jika ada
	// Misalnya, Anda dapat menggunakan metode dari productRepository
	return nil
}
