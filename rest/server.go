package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"log"
	"net/http"
)

type Server struct {
	productHandler *product.Handler
	userHandler *user.Handler
}

func NewServer(productHandler *product.Handler, userHandler *user.Handler) *Server {
	return &Server{
		productHandler: productHandler,
		userHandler: userHandler,
	}
}

func (s *Server) StartServer(cnf config.Config) {

	manager := middleware.NewManager()

	manager.Use(
		middleware.Logger,
		middleware.CorsMiddleware,
		middleware.PreflightMiddleware,
	)
	mux := http.NewServeMux() // Router

	wrappedMux := manager.WrappedMux(mux)

	// SetupRoutes(mux, manager)
	s.productHandler.RegisterRoutes(mux, manager)
	s.userHandler.RegisterRoutes(mux, manager)

	adr := ":" + cnf.HttpPort

	log.Println("Starting server on :9090")

	log.Fatal(http.ListenAndServe(adr, wrappedMux))
}
