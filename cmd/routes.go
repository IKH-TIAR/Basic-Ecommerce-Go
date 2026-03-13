package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	// Global middleware
	mux.Handle("GET /products", manager.Chain(
		http.HandlerFunc(handlers.GetProductsHandler),
	)) // Route get products

	mux.Handle("POST /products", manager.Chain(
		http.HandlerFunc(handlers.CreateProduct),
	)) // Route to create products

	mux.Handle("GET /products/{id}", manager.Chain(
		http.HandlerFunc(handlers.GetProductByID),
	)) // Route to get a product by ID
}
