package routes

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func SetRoutes(mux *http.ServeMux, mm *middleware.MiddlewareManager) {

	// // Define Product routes here
	// mux.Handle("GET /hellow", mm.Apply(http.HandlerFunc(product.HellowHandler)))

	// mux.Handle("GET /getProduct", mm.Apply(middleware.AuthenticationMiddleware(product.GetProduct)))
	// mux.Handle("POST /createProduct", mm.Apply(http.HandlerFunc(product.CreateProduct)))
	// mux.Handle("GET /getProduct/{id}", mm.Apply(http.HandlerFunc(product.GetProductById)))
	// mux.Handle("PUT /updatedProduct/{id}", mm.Apply(http.HandlerFunc(product.UpdateProductById)))
	// mux.Handle("DELETE /deleteProduct/{id}", mm.Apply(http.HandlerFunc(product.DeleteProductById)))
	// mux.Handle("GET /about", mm.AddMiddleware(http.HandlerFunc(product.AboutHandler), middleware.LoggerMiddleware))

	// // Define Users routes here
	// mux.Handle("POST /users/login", mm.Apply(http.HandlerFunc(user.LoginHandler)))
	// mux.Handle("GET /users", mm.Apply(http.HandlerFunc(user.GetUser)))
	// mux.Handle("POST /users", mm.Apply(http.HandlerFunc(user.CreateUser)))
	// mux.Handle("GET /users/{id}", mm.Apply(http.HandlerFunc(user.GetUserById)))
	// mux.Handle("PUT /users/{id}", mm.Apply(http.HandlerFunc(user.UpdateUser)))
	// mux.Handle("DELETE /users/{id}", mm.Apply(http.HandlerFunc(user.DeleteUser)))
}
