package repo

import (
	"ecommerce/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type LoginRepo interface {
	FindUserByEmail(email string) (*model.User, error)
	ValidateLogin(userLogin model.User) (bool, error)
}

type loginRepo struct {
	dbConnection *sqlx.DB
}

func NewLoginRepo(dbConnection *sqlx.DB) LoginRepo {
	return &loginRepo{
		dbConnection: dbConnection,
	}
}
func (l *loginRepo) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := l.dbConnection.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}
func (l *loginRepo) ValidateLogin(userLogin model.User) (bool, error) {
	user, err := l.FindUserByEmail(userLogin.Email)
	if err != nil {
		return false, err
	}

	if user.Password == userLogin.Password {
		return true, nil
	}

	return false, nil
}
