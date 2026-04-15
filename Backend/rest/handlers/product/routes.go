package product

import (
	"ecommerce/rest/middleware"
	"net/http"
)

// RegisterRoutes initializes the HTTP routes for the application.
func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	// Create a middleware manager and load the necessary middleware

	// Define routes with middleware
	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(h.GetProduct),
		),
	)
	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(h.CreateProduct),
			h.middlewares.AuthenticateJWT,
		),
	)
	mux.Handle(
		"GET /products/{productId}",
		manager.With(
			http.HandlerFunc(h.GetProductById),
		),
	)
	mux.Handle(
		"PUT /products/{productId}", //PUT = update more than one field, PATCH = update one field
		manager.With(
			http.HandlerFunc(h.UpdateProduct),
			h.middlewares.AuthenticateJWT,
		),
	)
	mux.Handle(
		"DELETE /products/{productId}",
		manager.With(
			http.HandlerFunc(h.DeleteProduct),
			h.middlewares.AuthenticateJWT,
		),
	)
}
