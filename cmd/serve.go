package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	dbConn, err := db.NewConnection()
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo(dbConn)

	middlewares := middleware.NewMiddleware(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(userRepo)

	server := rest.NewServer(productHandler, userHandler, cnf)

	server.StartServer()

}
