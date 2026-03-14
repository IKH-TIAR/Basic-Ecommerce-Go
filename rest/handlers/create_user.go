package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request){
	var newUser database.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil{
		utils.WriteError(w, http.StatusBadRequest, "Provide a valid Json")
		return
	}

	usr := newUser.Store()

	utils.WriteJSON(w, http.StatusCreated, usr)
}