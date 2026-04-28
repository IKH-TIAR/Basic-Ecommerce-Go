package user

import (
	userDomain "ecommerce/user"
)


type Handler struct {
	svc userDomain.UserService
}

func NewHandler(svc userDomain.UserService) *Handler {
	return &Handler{
		svc: svc,
	}
}