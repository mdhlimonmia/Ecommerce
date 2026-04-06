package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

// initRoutes initializes the HTTP routes for the application.
func initRoutes(mux *http.ServeMux) {

	// Create a middleware manager and load the necessary middleware
	manager := middleware.NewManager()
	manager.LoadMiddleware(middleware.Logger)

	// Define routes with middleware
	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProduct),
		),
	)
	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
			middleware.AuthMiddleware,
		),
	)
	mux.Handle(
		"GET /products/{productId}",
		manager.With(
			http.HandlerFunc(handlers.GetProductById),
		),
	)

}
