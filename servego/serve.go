package servego

import (
	"ecommerce/middleware"
	"ecommerce/routes"
	"fmt"
	"net/http"
)

func ServeGo() {

	setMiddleware := middleware.NewMiddlewareManager()

	setMiddleware.Use(middleware.LoggerMiddleware, middleware.CorsMiddleware, middleware.CorsAndPreflightMiddleware)

	mux := http.NewServeMux()

	routes.SetRoutes(mux, setMiddleware)

	muxWraped := setMiddleware.Apply(mux)

	fmt.Println("server running on :3000")

	err := http.ListenAndServe(":3000", muxWraped)

	if err != nil {
		fmt.Println(err, "error from server ")
	}
}
