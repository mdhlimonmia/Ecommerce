package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middleware"
	"net/http"
)

// initRoutes initializes the HTTP routes for the application.
func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	// Create a middleware manager and load the necessary middleware

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
			middleware.AuthenticateJWT,
		),
	)
	mux.Handle(
		"GET /products/{productId}",
		manager.With(
			http.HandlerFunc(handlers.GetProductById),
		),
	)
	mux.Handle(
		"PUT /products/{productId}", //PUT = update more than one field, PATCH = update one field
		manager.With(
			http.HandlerFunc(handlers.UpdateProduct),
			middleware.AuthenticateJWT,
		),
	)
	mux.Handle(
		"DELETE /products/{productId}",
		manager.With(
			http.HandlerFunc(handlers.DeleteProduct),
			middleware.AuthenticateJWT,
		),
	)
	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(handlers.Create_user),
		),
	)
	mux.Handle(
		"GET /users",
		manager.With(
			http.HandlerFunc(handlers.GetUsers),
			middleware.AuthenticateJWT,
		),
	)
	mux.Handle(
		"POST /users/login",
		manager.With(
			http.HandlerFunc(handlers.Login),
		),
	)

}
