package product

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	// Global middleware
	mux.Handle("GET /products", manager.Chain(
		http.HandlerFunc(h.GetProductsHandler),
	)) // Route get products

	mux.Handle("POST /products", manager.Chain(
		http.HandlerFunc(h.CreateProduct),
		middleware.AuthenticateJWT,
	)) // Route to create products

	mux.Handle("GET /products/{id}", manager.Chain(
		http.HandlerFunc(h.GetProduct),
	)) // Route to get a product by ID

	mux.Handle("PUT /products/{id}", manager.Chain(
		http.HandlerFunc(h.Update),
		middleware.AuthenticateJWT,
	)) // Route to update a product by ID

	mux.Handle("DELETE /products/{id}", manager.Chain(
		http.HandlerFunc(h.DeleteProduct),
		middleware.AuthenticateJWT,
	)) // Route to delete product by ID

}
