package product

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product,err2 := h.productRepo.Get(pID)
	if err2 != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to get product")
		return
	}
	if product != nil {
		utils.WriteJSON(w, http.StatusOK, product)
		return
	}
	
	utils.WriteError(w, http.StatusNotFound, "Product Not Found")
	
}
