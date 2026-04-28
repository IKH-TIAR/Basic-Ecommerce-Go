package user

import "ecommerce/domain"

type UserService interface {
	Create(user *domain.User) (*domain.User, error)
	Find(email, password string) (*domain.User, error)

}