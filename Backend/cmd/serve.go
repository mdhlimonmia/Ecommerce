package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"os"
)

func Serve() {

	dbCon, err := db.NewConnection()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		os.Exit(1)
	}

	cnf := config.GetConfig()
	m := middleware.NewMiddlewares(cnf)

	productRepo := repo.NewProductRepo()
	productHandler := product.NewHandler(m, productRepo)

	userRepo := repo.NewUserRepo(dbCon)
	userHandler := user.NewHandler(m, userRepo, cnf)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
