package repo

import "ecommerce/model"

type LoginRepo interface {
	FindUserByEmail(email string) model.User
	ValidateLogin(userLogin model.UserLogin) bool
}

type loginRepo struct {
	userList []model.User
}

func NewLoginRepo() LoginRepo {
	return &loginRepo{}
}
func (l *loginRepo) FindUserByEmail(email string) model.User {
	for _, user := range l.userList {
		if user.Email == email {	
			return user
		}
	}
	return model.User{}
}

func (l *loginRepo) ValidateLogin(userLogin model.UserLogin) bool {
	user := l.FindUserByEmail(userLogin.Email)
	if user.Id != 0 && user.Password == userLogin.Password {
		return true
	}
	return false
}