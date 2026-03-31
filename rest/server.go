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
	cnf *config.Config
	productHandler *product.Handler
	userHandler *user.Handler
}

func NewServer(
	productHandler *product.Handler, 
	userHandler *user.Handler,
	cnf *config.Config,
	) *Server {
	return &Server{
		cnf: cnf,
		productHandler: productHandler,
		userHandler: userHandler,
	}
}

func (s *Server) StartServer() {

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

	adr := ":" + s.cnf.HttpPort

	log.Println("Starting server on " + adr)

	log.Fatal(http.ListenAndServe(adr, wrappedMux))
}
