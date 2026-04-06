package cmd

import (
	"ecommerce/global_router"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()
	initRoutes(mux)

	fmt.Println("Server Running on :3080")

	//try to catch error
	err := http.ListenAndServe(":3080", global_router.GlobalRouter(mux))
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
