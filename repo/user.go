package repo

import "fmt"

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	Create(user *User) (*User, error)
	Find(email, password string) (*User, error)
}

type userRepo struct {
	userList []*User
}

func NewUserRepo() UserRepo {
	return &userRepo{
		userList: []*User{},
	}
}

func (r *userRepo) Create(user *User) (*User, error) {
	user.ID = len(r.userList) + 1
	r.userList = append(r.userList, user)
	return user, nil
}

func (r *userRepo) Find(email, password string ) (*User, error) {
	for _, user := range r.userList {
		if user.Email == email && user.Password == password {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}