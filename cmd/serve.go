package servego

import (
	"ecommerce/rest"
	"ecommerce/rest/controller/product"
	"ecommerce/rest/controller/user"
)

func ServeGo() {

	productHandler := product.NewHandler()

	userHandler := user.NewHandler()

	rest.NewServer(userHandler, productHandler).Start()
}
