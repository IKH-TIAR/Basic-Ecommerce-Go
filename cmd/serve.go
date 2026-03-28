package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
)

func Serve() {

	cnf := config.GetConfig()
	middlewares := middleware.NewMiddleware(cnf)

	productHandler := product.NewHandler(middlewares)

	userHandler := user.NewHandler()

	server := rest.NewServer(productHandler, userHandler, cnf)

	server.StartServer()

}
