package product

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct repo.Product

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

	createdProduct, err := h.productRepo.Create(newProduct)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Failed to create product")
		return
	}

	utils.WriteJSON(w, http.StatusCreated, createdProduct)

}
