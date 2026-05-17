package product

import "ecommerce/domain"

type Service struct {
	productRepo ProductRepo
}

func NewService(productRepo ProductRepo) *Service {
	return &Service{
		productRepo: productRepo,
	}
}

func (s *Service) Create(p domain.Product) (*domain.Product, error) {
	return s.productRepo.Create(p)
}

func (s *Service) Get(id int) (*domain.Product, error) {
	return s.productRepo.Get(id)
}

func (s *Service) List(limit, page int) ([]*domain.Product, error) {
	return s.productRepo.List(limit, page)
}

func (s *Service) TotalProducts() (int, error) {
	return s.productRepo.TotalProducts()
}

func (s *Service) Update(id int, p domain.Product) (*domain.Product, error) {
	return s.productRepo.Update(id, p)
}

func (s *Service) Delete(id int) error {
	return s.productRepo.Delete(id)
}
