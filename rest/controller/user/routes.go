package user

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) SetRoutesUser(mux *http.ServeMux, mm *middleware.MiddlewareManager) {

	// Define Users routes here
	mux.Handle("POST /users/login", mm.Apply(http.HandlerFunc(h.LoginHandler)))
	mux.Handle("GET /users", mm.Apply(http.HandlerFunc(h.GetUser)))
	mux.Handle("POST /users", mm.Apply(http.HandlerFunc(h.CreateUser)))
	mux.Handle("GET /users/{id}", mm.Apply(http.HandlerFunc(h.GetUserById)))
	mux.Handle("PUT /users/{id}", mm.Apply(http.HandlerFunc(h.UpdateUser)))
	mux.Handle("DELETE /users/{id}", mm.Apply(http.HandlerFunc(h.DeleteUser)))
}
