package product

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)


func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request){
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)

	if err!=nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res, err1 := h.productRepo.Delete(pID)
	if err1 != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	if res == "Deleted" {
		utils.WriteJSON(w, http.StatusOK, "Product Deleted")
		return
	}else {
		utils.WriteError(w, http.StatusNotFound, "Product Not Found")
	}

	
}