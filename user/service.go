package user

import "ecommerce/domain"

type userService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) UserService {
	return &userService{
		repo: repo,
	}
}


func (svc *userService) Create(user *domain.User) (*domain.User, error) {
	return svc.repo.Create(user)
}

func (svc *userService) Find(email, password string) (*domain.User, error) {
	return svc.repo.Find(email, password)
}
