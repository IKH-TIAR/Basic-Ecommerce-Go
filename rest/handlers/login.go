package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type ReqUsr struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var reqUsr ReqUsr

	if err := json.NewDecoder(r.Body).Decode(&reqUsr); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid Json")
		return
	}

	usr := database.Find(reqUsr.Email, reqUsr.Password)

	if usr == nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid User")
		return
	}

	utils.WriteJSON(w, http.StatusOK, usr)


}