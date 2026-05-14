package user

import (
	"ecommerce/domain"
	userHandler "ecommerce/rest/handlers/user"
)

type Service interface {
	userHandler.Service
}

type UserRepo interface {
	Create(u domain.User) (*domain.User, error)
	UserList() ([]*domain.User, error)
	FindUser(email string) bool
	AuthUser(email string, pass string) (*domain.User, error)
}
