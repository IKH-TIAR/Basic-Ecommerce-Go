package cmd

	import (
	"log"
	"net/http"
	"ecommerce/handlers"
	"ecommerce/middleware"

)

func StartServer() {
	mux := http.NewServeMux() // Router

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
	
	log.Println("Starting server on :9090")

	log.Fatal(http.ListenAndServe(":9090", middleware.CorsMiddleware(mux)))
}