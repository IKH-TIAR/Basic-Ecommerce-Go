package product

import "ecommerce/domain"

type ProductRepo interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(id int) (*domain.Product, error)
	List(page, limit int) ([]*domain.Product, error)
	Update(p domain.Product) (*domain.Product, error)
	Delete(id int) error
	Count() (int, error)
}

type ProductService interface {
	Create(p domain.Product) (*domain.Product, error)
	Get(id int) (*domain.Product, error)
	List(page, limit int) ([]*domain.Product, error)
	Update(p domain.Product) (*domain.Product, error)
	Count() (int, error)
	Delete(id int) error
}
