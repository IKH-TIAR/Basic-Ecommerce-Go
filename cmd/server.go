package cmd

	import (
	"log"
	"net/http"
	"ecommerce/middleware"

)

func StartServer() {
	mux := http.NewServeMux() // Router

	SetupRoutes(mux)
	
	log.Println("Starting server on :9090")

	log.Fatal(http.ListenAndServe(":9090", middleware.CorsMiddleware(mux)))
}