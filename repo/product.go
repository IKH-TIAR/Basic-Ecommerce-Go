package repo

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt	time.Time `json:"updated_at" db:"updated_at"`  
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id int) (*Product, error)
	List() ([]*Product, error)
	Update(p Product) (*Product, error)
	Delete(id int)  error
}

type productRepo struct {
	dbConn *sqlx.DB
}

func NewProductRepo(dbConn *sqlx.DB) ProductRepo {
	return &productRepo{
		dbConn: dbConn,
	}
}

func (r *productRepo) Create(p Product) (*Product, error) {
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

func (r *productRepo) Get(id int) (*Product, error) {
	query := `
	select id, title, description, price
	from products where id = $1
	`
	var product Product
	err := r.dbConn.Get(&product, query, id)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &product, nil

}

func (r *productRepo) List() ([]*Product, error) {
	query := `
	select * from products
	`
	var products []*Product
	err := r.dbConn.Select(&products, query)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return products, nil

}

func (r *productRepo) Update(p Product) (*Product, error) {

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
