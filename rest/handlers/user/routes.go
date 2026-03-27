package user

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("POST /users", manager.Chain(
		http.HandlerFunc(h.CreateUser),
	)) // Route to Register User

	mux.Handle("POST /users/login", manager.Chain(
		http.HandlerFunc(h.Login),
	)) // Route to Login User
}
