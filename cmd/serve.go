package servego

import (
	"ecommerce/rest"
	"ecommerce/rest/controller/product"
	"ecommerce/rest/controller/review"
	"ecommerce/rest/controller/user"
)

func ServeGo() {

	productHandler := product.NewHandler()

	userHandler := user.NewHandler()

	reviewHandler := review.NewHandler()

	rest.NewServer(userHandler, productHandler, reviewHandler).Start()
}
