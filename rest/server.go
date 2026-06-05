package rest

import (
	"ecommerce/config"
	"ecommerce/rest/controller/product"
	"ecommerce/rest/controller/review"
	"ecommerce/rest/controller/user"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
)

type Server struct {
	userHandler    *user.Handler
	productHandler *product.Handler
	reviewHandler  *review.Handler
}

func NewServer(userHandler *user.Handler,
	productHandler *product.Handler,
	reviewHandler *review.Handler) *Server {
	return &Server{
		userHandler:    userHandler,
		productHandler: productHandler,
		reviewHandler:  reviewHandler,
	}
}

func (s *Server) Start() {
	cfg, errs := config.Load()
	if errs != nil {
		panic(errs)
	}

	fmt.Println(cfg.Port)
	fmt.Println(cfg.JWTSecret)

	setMiddleware := middleware.NewMiddlewareManager()

	setMiddleware.Use(middleware.LoggerMiddleware, middleware.CorsMiddleware, middleware.CorsAndPreflightMiddleware)

	mux := http.NewServeMux()

	s.userHandler.SetRoutesUser(mux, setMiddleware)
	s.productHandler.SetRoutesProduct(cfg, mux, setMiddleware)
	s.reviewHandler.ReviewRoutes(mux, setMiddleware)

	// routes.SetRoutes(mux, setMiddleware)

	muxWraped := setMiddleware.Apply(mux)

	fmt.Println("server running on " + cfg.Port)

	err := http.ListenAndServe(":"+cfg.Port, muxWraped)

	if err != nil {
		fmt.Println(err, "error from server ")
	}
}
