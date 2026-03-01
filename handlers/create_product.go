package handlers

import (
	"encoding/json"
	"net/http"
	"ecommerce/database"
	"ecommerce/utils"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&newProduct); err != nil {
		http.Error(w, "Please Provide a valid json", http.StatusBadRequest)
		return
	}

	// validate
	if err := newProduct.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newProduct.ID = len(database.ProductList) + 1

	database.ProductList = append(database.ProductList, newProduct)

	utils.WriteJSON(w, http.StatusCreated, newProduct)

}