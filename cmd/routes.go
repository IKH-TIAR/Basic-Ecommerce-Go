package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) {
	mux.Handle("GET /products", middleware.Chain(
		http.HandlerFunc(handlers.GetProductsHandler),
		middleware.LoggerMiddleware,
	)) // Route get products

	mux.Handle("POST /products", middleware.Chain(
		http.HandlerFunc(handlers.CreateProduct),
		middleware.LoggerMiddleware,
	)) // Route to create products

	mux.Handle("GET /products/{id}", middleware.Chain(
		http.HandlerFunc(handlers.GetProductByID),
		middleware.LoggerMiddleware,
	)) // Route to get a product by ID
}
