package product

import (
	"ecommerce/config"
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) SetRoutesProduct(cfg *config.Config, mux *http.ServeMux, mm *middleware.MiddlewareManager) {

	// Define Product routes here
	mux.Handle("GET /hellow", mm.Apply(http.HandlerFunc(h.HellowHandler)))

	mux.Handle("GET /getProduct", mm.Apply(middleware.AuthenticationMiddleware(cfg.JWTSecret, h.GetProduct)))
	mux.Handle("POST /createProduct", mm.Apply(http.HandlerFunc(h.CreateProduct)))
	mux.Handle("GET /getProduct/{id}", mm.Apply(http.HandlerFunc(h.GetProductById)))
	mux.Handle("PUT /updatedProduct/{id}", mm.Apply(http.HandlerFunc(h.UpdateProductById)))
	mux.Handle("DELETE /deleteProduct/{id}", mm.Apply(http.HandlerFunc(h.DeleteProductById)))
	mux.Handle("GET /about", mm.AddMiddleware(http.HandlerFunc(h.AboutHandler), middleware.LoggerMiddleware))

}
