package product

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&newProduct); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Please Provide a Valid Json")
		return
	}

	// validate
	if err := newProduct.Validate(); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	createdProduct := database.Store(newProduct)

	utils.WriteJSON(w, http.StatusCreated, createdProduct)

}
