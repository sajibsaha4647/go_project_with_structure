package rest

import (
	"ecommerce/rest/middleware"
	"ecommerce/rest/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Start() {
	errs := godotenv.Load()
	if errs != nil {
		panic(errs)
	}

	fmt.Println(os.Getenv("PORT"))
	fmt.Println(os.Getenv("DB_NAME"))

	setMiddleware := middleware.NewMiddlewareManager()

	setMiddleware.Use(middleware.LoggerMiddleware, middleware.CorsMiddleware, middleware.CorsAndPreflightMiddleware)

	mux := http.NewServeMux()

	routes.SetRoutes(mux, setMiddleware)

	muxWraped := setMiddleware.Apply(mux)

	fmt.Println("server running on " + os.Getenv("PORT"))

	err := http.ListenAndServe(":"+os.Getenv("PORT"), muxWraped)

	if err != nil {
		fmt.Println(err, "error from server ")
	}
}
