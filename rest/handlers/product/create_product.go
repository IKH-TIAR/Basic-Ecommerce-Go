package product

import (
	"ecommerce/domain"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct domain.Product

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&newProduct); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Please Provide a Valid Json")
		return
	}

	// // validate
	// if err := newProduct.Validate(); err != nil {
	// 	utils.WriteError(w, http.StatusBadRequest, err.Error())
	// 	return
	// }

	createdProduct, err := h.svc.Create(newProduct)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to create product")
		return
	}

	utils.WriteJSON(w, http.StatusCreated, createdProduct)

}
