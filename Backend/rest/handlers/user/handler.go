package user

import (
	"ecommerce/config"
	"ecommerce/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	svc         Service
	cnf         *config.Config
}

func NewHandler(
	middlewares *middleware.Middlewares,
	svc Service,
	cnf *config.Config,
) *Handler {
	return &Handler{
		middlewares: middlewares,
		svc:         svc,
		cnf:         cnf,
	}
}
