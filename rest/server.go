package rest

import (
	"ecommerce/rest/controller/product"
	"ecommerce/rest/controller/user"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Server struct {
	userHandler    *user.Handler
	productHandler *product.Handler
}

func NewServer(userHandler *user.Handler,
	productHandler *product.Handler) *Server {
	return &Server{
		userHandler:    userHandler,
		productHandler: productHandler,
	}
}

func (s *Server) Start() {
	errs := godotenv.Load()
	if errs != nil {
		panic(errs)
	}

	fmt.Println(os.Getenv("PORT"))
	fmt.Println(os.Getenv("DB_NAME"))

	setMiddleware := middleware.NewMiddlewareManager()

	setMiddleware.Use(middleware.LoggerMiddleware, middleware.CorsMiddleware, middleware.CorsAndPreflightMiddleware)

	mux := http.NewServeMux()

	s.userHandler.SetRoutesUser(mux, setMiddleware)
	s.productHandler.SetRoutesProduct(mux, setMiddleware)

	// routes.SetRoutes(mux, setMiddleware)

	muxWraped := setMiddleware.Apply(mux)

	fmt.Println("server running on " + os.Getenv("PORT"))

	err := http.ListenAndServe(":"+os.Getenv("PORT"), muxWraped)

	if err != nil {
		fmt.Println(err, "error from server ")
	}
}
