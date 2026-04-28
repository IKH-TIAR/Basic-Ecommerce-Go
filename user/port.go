package user

import (
	"ecommerce/domain"
	userHandler "ecommerce/rest/handlers/user"
)


type UserService interface {
	userHandler.UserService
}

type UserRepo interface {
	Create(user *domain.User) (*domain.User, error)
	Find(email, password string) (*domain.User, error)
}
