package handlers

import (
	"net/http"
	services "simple-checkout-app/internal/service"
)

type ProductHandler interface {
	AddToCart(w http.ResponseWriter, r *http.Request)
	GetCart(w http.ResponseWriter, r *http.Request)
	RemoveFromCart(w http.ResponseWriter, r *http.Request)
}

type productHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) ProductHandler {
	return &productHandler{productService}
}

func (h *productHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	// Implementasi untuk menambahkan produk ke keranjang
}

func (h *productHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	// Implementasi untuk mendapatkan daftar produk dalam keranjang
}

func (h *productHandler) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	// Implementasi untuk menghapus produk dari keranjang
}
