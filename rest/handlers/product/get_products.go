package product

import (
	"net/http"
	"ecommerce/database"
	"ecommerce/utils"
)

func (h *Handler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	
	utils.WriteJSON(w, http.StatusOK, database.List())
}