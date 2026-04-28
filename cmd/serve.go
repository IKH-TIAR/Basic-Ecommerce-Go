package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	productHandler "ecommerce/rest/handlers/product"
	userHandler "ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"ecommerce/user"
)

func Serve() {
	cnf := config.GetConfig()

	dbConn, err := db.NewConnection(cnf.Database)
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// Initialize repositories
	productRepo := repo.NewProductRepo(dbConn)
	userRepo := repo.NewUserRepo(dbConn)

	// Initialize domain
	userService := user.NewUserService(userRepo)

	middlewares := middleware.NewMiddleware(cnf)

	productHandler := productHandler.NewHandler(middlewares, productRepo)
	userHandler := userHandler.NewHandler(userService)

	server := rest.NewServer(productHandler, userHandler, cnf)

	server.StartServer()

}
