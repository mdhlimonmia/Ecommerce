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
	m := middleware.NewMiddlewares(cnf)

	productRepo := repo.NewProductRepo()
	productHandler := product.NewHandler(m, productRepo)

	userRepo := repo.NewUserRepo()
	userHandler := user.NewHandler(m, userRepo, cnf)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
