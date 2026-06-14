package user

import "ecommerce/repo"

type Handler struct {
	userRepo repo.UserRepo
	loginRepo repo.LoginRepo
}

func NewHandler(userRepo repo.UserRepo, loginRepo repo.LoginRepo) *Handler {
	return &Handler{
		userRepo: userRepo,
		loginRepo: loginRepo,
	}
}
