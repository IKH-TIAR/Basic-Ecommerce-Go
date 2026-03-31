package product

import (
	"net/http"
	"ecommerce/utils"
)

func (h *Handler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := h.productRepo.List()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to get products")
		return
	}
	utils.WriteJSON(w, http.StatusOK, products)
}