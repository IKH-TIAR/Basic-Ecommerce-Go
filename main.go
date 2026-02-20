package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func testHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Test Route")
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

var productlist []Product

func getProductsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Please Give a valid Request Method", 404)
		return
	}

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(productlist); err != nil {
		http.Error(w, "Error encoding JSON", 500)
		return
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", testHandler)

	mux.HandleFunc("/products", getProductsHandler)

	log.Println("Starting server on :9090")

	log.Fatal(http.ListenAndServe(":9090", mux))

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
	productlist = append(productlist, prd1, prd2, prd3, prd4)
}
