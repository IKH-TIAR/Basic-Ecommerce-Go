package database


type User struct{
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	IsShopOwner bool `json:"is_shop_owner"`
}

var userList []User

func (u User) Store() User {
	uID := len(userList) + 1
	u.ID = uID
	userList = append(userList, u)
	return u
}

func Find(email, password string) *User {
	for _, user := range userList{
		if user.Email == email && user.Password == password {
			return &user
		}
	}
	return nil
}