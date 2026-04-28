package repo

import (
	"database/sql"
	"ecommerce/user"
	"ecommerce/domain"
	"log"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	dbConn *sqlx.DB
}

func NewUserRepo(dbConn *sqlx.DB) UserRepo {
	return &userRepo{
		dbConn: dbConn,
	}
}

func (r *userRepo) Create(user *domain.User) (*domain.User, error) {
	query := `INSERT INTO users (
	first_name, 
	last_name, 
	email, 
	password, 
	is_shop_owner
	) VALUES 
	(
	:first_name, 
	:last_name,
	:email,
	:password,
	:is_shop_owner
	) RETURNING id`

	var userID int

	rows, err := r.dbConn.NamedQuery(query, user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if rows.Next(){
		rows.Scan(&userID)
	}
	user.ID = userID

	return user, nil

}

func (r *userRepo) Find(email, password string ) (*domain.User, error) {
	var user domain.User
	query := `
	SELECT id, first_name, last_name, email, password, is_shop_owner
	FROM users
	WHERE email = $1 AND password = $2
	LIMIT 1
	`

	err := r.dbConn.Get(&user, query, email, password)

	if err != nil {
		if err == sql.ErrNoRows{
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}