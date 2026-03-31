package user

import (
	"ecommerce/config"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

type ReqUsr struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqUsr ReqUsr

	if err := json.NewDecoder(r.Body).Decode(&reqUsr); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid Json")
		return
	}

	usr, err1 := h.userRepo.Find(reqUsr.Email, reqUsr.Password)

	if err1 != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid User")
		return
	}

	if usr == nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid User")
		return
	}
	cnf := config.GetConfig()

	AccessToken := utils.CreateJWT(cnf.Secret, utils.Payload{
		Sub:         usr.ID,
		FirstName:   usr.FirstName,
		LastName:    usr.LastName,
		Email:       usr.Email,
		IsShopOwner: usr.IsShopOwner,
	})

	utils.WriteJSON(w, http.StatusOK, AccessToken)

}
