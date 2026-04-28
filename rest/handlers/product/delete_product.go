package product

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err1 := h.svc.Delete(pID)
	if err1 != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	utils.WriteJSON(w, http.StatusOK, "Product Deleted")

}
