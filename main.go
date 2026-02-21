package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
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

func createProduc(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions{
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Please Give Post Request", 400)
		return
	}

	var newProduct Product

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&newProduct); err != nil {
		http.Error(w, "Please Provide a valid json", http.StatusBadRequest)
		log.Println(err)
		return
	}

	if strings.TrimSpace(newProduct.Description) == "" || newProduct.Price < 0.00 || strings.TrimSpace(newProduct.Title) == "" {
		http.Error(w, "Fields Cannot be Empty", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(productlist) + 1

	productlist = append(productlist, newProduct)

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(newProduct); err != nil {
		http.Error(w, "Error Encoding", http.StatusInternalServerError)
		return
	}



}

func main() {
	mux := http.NewServeMux() // Router

	mux.HandleFunc("/", testHandler) // Route

	mux.HandleFunc("/products", getProductsHandler) // Route get products

	mux.HandleFunc("/create", createProduc) // Route to create products

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
