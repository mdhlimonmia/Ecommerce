package product

import "ecommerce/domain"

type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(id int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Update(id int, p domain.Product) (*domain.Product, error)
	Delete(id int) error
}
