package product

import "ecommerce/domain"


type productService struct {
	repo ProductRepo
}

func NewProductService (repo ProductRepo) ProductService {
	return &productService{
		repo: repo,
	}
}


func (s *productService) Create(p domain.Product) (*domain.Product, error) {
	product, err := s.repo.Create(p)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) Get(id int) (*domain.Product, error) {
	product, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) List() ([]*domain.Product, error) {
	products, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *productService) Update(p domain.Product) (*domain.Product, error) {
	product, err := s.repo.Update(p)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

