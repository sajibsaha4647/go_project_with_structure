package rest

import (
	"ecommerce/config"
	"ecommerce/rest/controller/product"
	"ecommerce/rest/controller/user"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
)

type Server struct {
	userHandler    *user.Handler
	productHandler *product.Handler
}

func NewServer(userHandler *user.Handler, productHandler *product.Handler) *Server {
	return &Server{
		userHandler:    userHandler,
		productHandler: productHandler,
	}
}

func (s *Server) Start(cnf *config.Config) {
	setMiddleware := middleware.NewMiddlewareManager()
	setMiddleware.Use(middleware.LoggerMiddleware, middleware.CorsMiddleware, middleware.CorsAndPreflightMiddleware)

	mux := http.NewServeMux()

	s.userHandler.SetRoutesUser(mux, setMiddleware)
	s.productHandler.SetRoutesProduct(cnf, mux, setMiddleware)

	muxWraped := setMiddleware.Apply(mux)

	fmt.Println("server running on " + cnf.Port)

	if err := http.ListenAndServe(":"+cnf.Port, muxWraped); err != nil {
		fmt.Println(err, "error from server")
	}
}
