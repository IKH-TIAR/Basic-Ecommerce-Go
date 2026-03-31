package repo

import "slices"

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id int) (*Product, error)
	List() ([]*Product, error)
	Update(p Product) (*Product, error)
	Delete(id int) (string, error)
}

type productRepo struct {
	productList []*Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateIntialProduct(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(id int) (*Product, error) {

	for _, product := range r.productList {
		if product.ID == id {
			return product, nil
		}
	}
	return nil, nil

}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil

}

func (r *productRepo) Update(p Product) (*Product, error) {
	for idx, product := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &p
		}
	}
	return &p, nil
}

func (r *productRepo) Delete(id int) (string, error) {
	for idx, product := range r.productList {
		if id == product.ID {
			r.productList = slices.Delete(r.productList, idx, idx+1)
			return "Deleted", nil
		}
	}
	return "", nil
}

func generateIntialProduct(r *productRepo) {
	prd1 := &Product{
		ID:          1,
		Title:       "Mango",
		Description: "This is a mango, we like to eat mango",
		Price:       45.44,
	}
	prd2 := &Product{
		ID:          2,
		Title:       "orange",
		Description: "This is a orange, we like to eat orange",
		Price:       45.44,
	}

	prd3 := &Product{
		ID:          3,
		Title:       " banana",
		Description: "This is a banana, we don't like to eat banana",
		Price:       45.44,
	}
	prd4 := &Product{
		ID:          4,
		Title:       "Not banana",
		Description: "This is a not banana, we don't like to eat not banana",
		Price:       45.44,
	}
	r.productList = append(r.productList, prd1, prd2, prd3, prd4)
}
