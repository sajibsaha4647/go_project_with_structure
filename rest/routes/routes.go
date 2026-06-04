package routes

import (
	"ecommerce/rest/controller"
	"ecommerce/rest/middleware"
	"net/http"
)

func SetRoutes(mux *http.ServeMux, mm *middleware.MiddlewareManager) {

	// Define Product routes here
	mux.Handle("GET /hellow", mm.Apply(http.HandlerFunc(controller.HellowHandler)))
	
	mux.Handle("GET /getProduct", mm.Apply(middleware.AuthenticationMiddleware(controller.GetProduct)))
	mux.Handle("POST /createProduct", mm.Apply(http.HandlerFunc(controller.CreateProduct)))
	mux.Handle("GET /getProduct/{id}", mm.Apply(http.HandlerFunc(controller.GetProductById)))
	mux.Handle("PUT /updatedProduct/{id}", mm.Apply(http.HandlerFunc(controller.UpdateProductById)))
	mux.Handle("DELETE /deleteProduct/{id}", mm.Apply(http.HandlerFunc(controller.DeleteProductById)))
	mux.Handle("GET /about", mm.AddMiddleware(http.HandlerFunc(controller.AboutHandler), middleware.LoggerMiddleware))

	// Define Users routes here
	mux.Handle("POST /users/login", mm.Apply(http.HandlerFunc(controller.LoginHandler)))
	mux.Handle("GET /users", mm.Apply(http.HandlerFunc(controller.GetUser)))
	mux.Handle("POST /users", mm.Apply(http.HandlerFunc(controller.CreateUser)))
	mux.Handle("GET /users/{id}", mm.Apply(http.HandlerFunc(controller.GetUserById)))
	mux.Handle("PUT /users/{id}", mm.Apply(http.HandlerFunc(controller.UpdateUser)))
	mux.Handle("DELETE /users/{id}", mm.Apply(http.HandlerFunc(controller.DeleteUser)))
}
