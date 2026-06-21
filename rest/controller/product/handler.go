package product

import (
	productService "ecommerce/product"
)

type Handler struct {
	service productService.Service
}

func NewHandler(repo productService.Service) *Handler {
	return &Handler{
		service: repo,
	}
}
