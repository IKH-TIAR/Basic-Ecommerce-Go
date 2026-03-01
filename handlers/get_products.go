package handlers

import (
	"net/http"
	"ecommerce/database"
	"ecommerce/utils"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	
	utils.WriteJSON(w, http.StatusOK, database.ProductList)
}