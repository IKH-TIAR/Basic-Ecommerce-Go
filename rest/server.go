package rest

import (
	"ecommerce/config"
	"ecommerce/rest/middleware"
	"log"
	"net/http"
)

func StartServer(cnf config.Config) {

		manager := middleware.NewManager()

	manager.Use(
		middleware.Logger,
		middleware.CorsMiddleware,
		middleware.PreflightMiddleware,
	)
	mux := http.NewServeMux() // Router

	wrappedMux := manager.WrappedMux(mux)

	SetupRoutes(mux, manager)

	adr := ":" + cnf.HttpPort

	log.Println("Starting server on :9090")

	log.Fatal(http.ListenAndServe(adr, wrappedMux))
}