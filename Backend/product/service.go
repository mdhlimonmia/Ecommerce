package product

import "ecommerce/domain"

type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(id int) (*domain.Product, error)
	List(limit, page int) ([]*domain.Product, error)
	TotalProducts() (int, error)
	Update(id int, p domain.Product) (*domain.Product, error)
	Delete(id int) error
}
