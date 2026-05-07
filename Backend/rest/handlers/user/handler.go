package user

import (
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/rest/middleware"
)

type Handler struct {
	middlewares *middleware.Middlewares
	userRepo    repo.UserRepo
	cnf         *config.Config
}

func NewHandler(
	middlewares *middleware.Middlewares,
	userRepo repo.UserRepo,
	cnf *config.Config,
) *Handler {
	return &Handler{
		middlewares: middlewares,
		userRepo:    userRepo,
		cnf:         cnf,
	}
}
