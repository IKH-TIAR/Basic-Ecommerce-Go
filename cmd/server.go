package cmd

	import (
	"log"
	"net/http"
	"ecommerce/handlers"
	"ecommerce/middleware"

)

func StartServer() {
	mux := http.NewServeMux() // Router

	mux.Handle("GET /products", middleware.LoggerMiddleware(http.HandlerFunc(handlers.GetProductsHandler))) // Route get products

	mux.Handle("POST /products", middleware.LoggerMiddleware(http.HandlerFunc(handlers.CreateProduct))) // Route to create products

	mux.Handle("GET /products/{id}", middleware.LoggerMiddleware(http.HandlerFunc(handlers.GetProductByID))) // Route to get a product by ID
	
	log.Println("Starting server on :9090")

	log.Fatal(http.ListenAndServe(":9090", middleware.CorsMiddleware(mux)))
}