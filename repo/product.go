package repo

import (
	"ecommerce/domain"
	"ecommerce/product"
	"log"

	"github.com/jmoiron/sqlx"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	dbConn *sqlx.DB
}

func NewProductRepo(dbConn *sqlx.DB) ProductRepo {
	return &productRepo{
		dbConn: dbConn,
	}
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
	query := `
	insert into products (
	title,
	description,
	price
	) values (
	$1,
	$2,
	$3
	) 
	returning id
	`
	err := r.dbConn.QueryRow(query, p.Title, p.Description, p.Price).Scan(&p.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &p, nil
}

func (r *productRepo) Get(id int) (*domain.Product, error) {
	query := `
	select id, title, description, price
	from products where id = $1
	`
	var product domain.Product
	err := r.dbConn.Get(&product, query, id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &product, nil

}

func (r *productRepo) List() ([]*domain.Product, error) {
	query := `
	select * from products
	`
	var products []*domain.Product
	err := r.dbConn.Select(&products, query)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return products, nil

}

func (r *productRepo) Update(p domain.Product) (*domain.Product, error) {

	query := `
	update products set
	title = $1,
	description = $2,
	price = $3
	where id = $4
	`

	row := r.dbConn.QueryRow(query, p.Title, p.Description, p.Price, p.ID)

	err := row.Err()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &p, nil


}

func (r *productRepo) Delete(id int)  error {
	query := `
	delete from products where id = $1
	`

	_, err := r.dbConn.Exec(query, id)

	if err != nil {
		return err
	}

	return nil

}
