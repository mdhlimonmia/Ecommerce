package user

import (
	"ecommerce/rest/middleware"
	"net/http"
)

// RegisterRoutes initializes the HTTP routes for the application.
func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"POST /users",
		manager.With(
			http.HandlerFunc(h.Create_user),
		),
	)
	mux.Handle(
		"GET /users",
		manager.With(
			http.HandlerFunc(h.GetUsers),
			h.middlewares.AuthenticateJWT,
		),
	)
	mux.Handle(
		"POST /users/login",
		manager.With(
			http.HandlerFunc(h.Login),
		),
	)

}
