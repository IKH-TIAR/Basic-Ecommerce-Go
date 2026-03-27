package product

import (
	"ecommerce/database"
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

	res := database.Delete(pID)

	if res == "deleted"{
		utils.WriteJSON(w, http.StatusOK, "Product Deleted")
		return
	}

	utils.WriteError(w, http.StatusNotFound, "Product Not Found")


}