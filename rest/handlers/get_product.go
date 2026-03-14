package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product := database.Get(pID)
	if product != nil {
		utils.WriteJSON(w, http.StatusOK, product)
		return
	}
	
	utils.WriteError(w, http.StatusNotFound, "Product Not Found")
	
}
