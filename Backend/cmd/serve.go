package cmd

import (
	"ecommerce/config"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	config := config.GetConfig()

	manager := middleware.NewManager()

	manager.LoadMiddleware(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	// Initialize routes
	initRoutes(mux, manager)

	port := ":" + config.HttpPort
	fmt.Println("Server Running on:", port)

	//try to catch error
	err := http.ListenAndServe(port, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
