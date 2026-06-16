package user

import userservice "ecommerce/user"

type Handler struct {
	service userservice.Service
}

func NewHandler(service userservice.Service) *Handler {
	return &Handler{service: service}
}
