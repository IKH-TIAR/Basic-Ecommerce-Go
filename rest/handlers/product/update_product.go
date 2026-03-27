package product

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) Update(w http.ResponseWriter, r *http.Request){
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var newProduct database.Product

	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	newProduct.ID = pID

	database.Update(newProduct)

	utils.WriteJSON(w, http.StatusOK, "Product Updated")
} 