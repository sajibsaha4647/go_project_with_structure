package product

import (
	"ecommerce/repo"
)

type Handler struct {
	repo repo.ProductRepo
}

func NewHandler(repo repo.ProductRepo) *Handler {
	return &Handler{
		repo: repo,
	}
}
