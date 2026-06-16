package user

import "ecommerce/domain"

type Repository interface {
	Store(domain.User) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
	GetUserById(id int) *domain.User
	UpdateUserById(id int, u domain.User) *domain.User
	DeleteUserById(id int) bool
	FindUserByEmail(email string) *domain.User
}
