package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
productID := r.PathValue("id")

pID, err := strconv.Atoi(productID)
if err != nil {
	http.Error(w, "Invalid product ID", http.StatusBadRequest)
	return
}
for _, product := range database.ProductList {
	if product.ID == pID {
		utils.WriteJSON(w, http.StatusOK, product)
		return
	}
	
}
http.Error(w, "Product not found", http.StatusNotFound)
}