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
	user, err :=svc.repo.Create(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (svc *userService) Find(email, password string) (*domain.User, error) {
	user, err := svc.repo.Find(email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
