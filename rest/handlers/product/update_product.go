package product

import (
	"ecommerce/repo"
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

	var newProduct repo.Product

	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	newProduct.ID = pID

	prd, err1 := h.productRepo.Update(newProduct)
	if err1 != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to update product")
		return
	}

	utils.WriteJSON(w, http.StatusOK, prd)
} 