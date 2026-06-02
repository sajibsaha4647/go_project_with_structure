package routes

import (
	controller "ecommerce/Controller"
	"ecommerce/middleware"
	"net/http"
)

func SetRoutes(mux *http.ServeMux, mm *middleware.MiddlewareManager) {
	mux.Handle("GET /hellow", mm.Apply(http.HandlerFunc(controller.HellowHandler)))
	mux.Handle("GET /getProduct", mm.Apply(http.HandlerFunc(controller.GetProduct)))
	mux.Handle("POST /createProduct", mm.Apply(http.HandlerFunc(controller.CreateProduct)))
	mux.Handle("GET /getProduct/{id}", mm.Apply(http.HandlerFunc(controller.GetProductById)))
	mux.Handle("GET /about", mm.AddMiddleware(http.HandlerFunc(controller.AboutHandler), middleware.LoggerMiddleware))
}
