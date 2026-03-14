package database

import (
	"errors"
	"slices"
	"strings"
)

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}


func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(id int) *Product {
	for _, product := range productList{
		if product.ID == id{
			return &product
		}
	}
	return nil
}

func Update(p Product){
	for idx, product := range productList{
		if p.ID == product.ID{
			productList[idx] = p
		}
	}
}

func Delete(id int) string {
	for idx, product := range productList{
		if id == product.ID{
			productList = slices.Delete(productList, idx, idx+1)
			return "deleted"
		}
	}

	return "not_found"
}


func (p Product) Validate() error {

	if strings.TrimSpace(p.Description) == "" || p.Price < 0.00 || strings.TrimSpace(p.Title) == "" {
		return errors.New("Fields Can Not Be Empty")
	}

	return nil
}



func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Mango",
		Description: "This is a mango, we like to eat mango",
		Price:       45.44,
	}
	prd2 := Product{
		ID:          2,
		Title:       "orange",
		Description: "This is a orange, we like to eat orange",
		Price:       45.44,
	}

	prd3 := Product{
		ID:          3,
		Title:       " banana",
		Description: "This is a banana, we don't like to eat banana",
		Price:       45.44,
	}
	prd4 := Product{
		ID:          4,
		Title:       "Not banana",
		Description: "This is a not banana, we don't like to eat not banana",
		Price:       45.44,
	}
	productList = append(productList, prd1, prd2, prd3, prd4)
	
}