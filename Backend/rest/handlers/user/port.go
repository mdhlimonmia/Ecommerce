package user

import "ecommerce/domain"

type Service interface {
	Create(u domain.User) (*domain.User, error)
	UserList() ([]*domain.User, error)
	FindUser(email string) bool
	AuthUser(email string, pass string) (*domain.User, error)
}
