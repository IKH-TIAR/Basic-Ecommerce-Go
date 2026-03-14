package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middleware"
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
		http.HandlerFunc(handlers.GetProduct),
	)) // Route to get a product by ID

	mux.Handle("PUT /products/{id}", manager.Chain(
		http.HandlerFunc(handlers.Update),
	)) // Route to update a product by ID

	mux.Handle("DELETE /products/{id}", manager.Chain(
		http.HandlerFunc(handlers.DeleteProduct),
	)) // Route to delete product by ID

	mux.Handle("POST /users", manager.Chain(
		http.HandlerFunc(handlers.CreateUser),
	)) // Route to Register User

	mux.Handle("POST /users/login", manager.Chain(
		http.HandlerFunc(handlers.Login),
	)) // Route to Login User
}
