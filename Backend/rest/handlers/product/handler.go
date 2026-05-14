package product

import (
	"ecommerce/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	svc         Service
}

func NewHandler(
	middlewares *middleware.Middlewares,
	service Service,
) *Handler {
	return &Handler{
		middlewares: middlewares,
		svc:         service,
	}
}
