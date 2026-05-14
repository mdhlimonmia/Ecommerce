package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/product"
	"ecommerce/repo"
	"ecommerce/rest"
	product_handler "ecommerce/rest/handlers/product"
	user_handler "ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"ecommerce/user"
	"fmt"
	"os"
)

func Serve() {

	dbCon, err := db.NewConnection(config.GetDbConfig())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println("Error migrating database:", err)
		os.Exit(1)
	}

	cnf := config.GetConfig()
	m := middleware.NewMiddlewares(cnf)

	//repo
	usrRepo := repo.NewUserRepo(dbCon)
	prodRepo := repo.NewProductRepo(dbCon)

	//service
	userSvc := user.NewService(usrRepo)
	prodSvc := product.NewService(prodRepo)

	userHandler := user_handler.NewHandler(m, userSvc, cnf)
	productHandler := product_handler.NewHandler(m, prodSvc)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
