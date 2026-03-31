package cmd

import (
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
)

func Serve() {

	cnf := config.GetConfig()
	middlewares := middleware.NewMiddleware(cnf)
	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()

	productHandler := product.NewHandler(middlewares, productRepo)

	userHandler := user.NewHandler(userRepo)

	server := rest.NewServer(productHandler, userHandler, cnf)

	server.StartServer()

}
