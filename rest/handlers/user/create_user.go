package user

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request){
	var newUser repo.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil{
		utils.WriteError(w, http.StatusBadRequest, "Provide a valid Json")
		return
	}

	usr, err := h.userRepo.Create(&newUser)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Server Error")
		return
	}

	utils.WriteJSON(w, http.StatusCreated, usr)
}