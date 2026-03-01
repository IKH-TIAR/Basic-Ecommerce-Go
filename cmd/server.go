package cmd

	import (
	"log"
	"net/http"
	"ecommerce/handlers"
	"ecommerce/middleware"
)

func StartServer() {
	mux := http.NewServeMux() // Router

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProductsHandler)) // Route get products

	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct)) // Route to create products

	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductByID)) // Route to get a product by ID
	
	log.Println("Starting server on :9090")

	log.Fatal(http.ListenAndServe(":9090", middleware.CorsMiddleware(mux)))
}