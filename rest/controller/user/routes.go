package user

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func SetRoutesUser(mux *http.ServeMux, mm *middleware.MiddlewareManager) {

	// Define Users routes here
	mux.Handle("POST /users/login", mm.Apply(http.HandlerFunc(LoginHandler)))
	mux.Handle("GET /users", mm.Apply(http.HandlerFunc(GetUser)))
	mux.Handle("POST /users", mm.Apply(http.HandlerFunc(CreateUser)))
	mux.Handle("GET /users/{id}", mm.Apply(http.HandlerFunc(GetUserById)))
	mux.Handle("PUT /users/{id}", mm.Apply(http.HandlerFunc(UpdateUser)))
	mux.Handle("DELETE /users/{id}", mm.Apply(http.HandlerFunc(DeleteUser)))
}
