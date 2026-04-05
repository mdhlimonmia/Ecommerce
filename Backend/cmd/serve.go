package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	mux.Handle("GET /product", http.HandlerFunc(handlers.GetProduct))
	mux.Handle("POST /product", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /product/{productId}", http.HandlerFunc(handlers.GetProductById))

	fmt.Println("Server Running on :3080")

	//try to catch error
	err := http.ListenAndServe(":3080", global_router.GlobalRouter(mux))
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
