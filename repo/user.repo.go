package repo

import "ecommerce/model"

type UserRepo interface {
	Store(model.User) []model.User
	GetAllUsers() []model.User
	GetUserById(id int) model.User
	UpdateUserById(id int, updatedUser model.User) model.User
	DeleteUserById(id int) bool
}

type userRepo struct {
	userList []model.User
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (u *userRepo) Store(user model.User) []model.User {
	if len(u.userList) == 0 {
		user.Id = 1
	} else {
		user.Id = u.userList[len(u.userList)-1].Id + 1
	}
	u.userList = append(u.userList, user)
	return u.userList
}
func (u *userRepo) GetAllUsers() []model.User {
	return u.userList
}
func (u *userRepo) GetUserById(id int) model.User {
	for _, user := range u.userList {
		if id == user.Id {
			return user
		}
	}
	return model.User{}
}
func (u *userRepo) UpdateUserById(id int, updatedUser model.User) model.User {
	for i, user := range u.userList {
		if id == user.Id {
			u.userList[i] = updatedUser
			return updatedUser
		}
	}
	return model.User{}
}
func (u *userRepo) DeleteUserById(id int) bool {
	for i, user := range u.userList {
		if id == user.Id {
			u.userList = append(u.userList[:i], u.userList[i+1:]...)
			return true
		}
	}
	return false
}

func (u *userRepo) FindUserByEmail(email string) model.User {
	for _, user := range u.userList {
		if email == user.Email {
			return user
		}
	}
	return model.User{}
}