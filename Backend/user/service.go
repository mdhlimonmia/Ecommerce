package user

import "ecommerce/domain"

type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) Create(u domain.User) (*domain.User, error) {
	return s.userRepo.Create(u)
}

func (s *service) UserList() ([]*domain.User, error) {
	return s.userRepo.UserList()
}

func (s *service) FindUser(email string) bool {
	return s.userRepo.FindUser(email)
}

func (s *service) AuthUser(email string, pass string) (*domain.User, error) {
	// Implementation for authenticating user
	return s.userRepo.AuthUser(email, pass)
}
