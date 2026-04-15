package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
)

func Serve() {
	config := config.GetConfig()
	m := middleware.NewMiddlewares(config)
	productHandler := product.NewHandler(m)
	userHandler := user.NewHandler(m)

	server := rest.NewServer(
		config,
		productHandler,
		userHandler,
	)
	server.Start()
}
