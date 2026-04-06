package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
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

	fmt.Println("Server Running on :3080")

	//try to catch error
	err := http.ListenAndServe(":3080", wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
