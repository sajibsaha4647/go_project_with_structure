package servego

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/product"
	"ecommerce/rest"
	pdCtrl "ecommerce/rest/controller/product"
	userctrl "ecommerce/rest/controller/user"
	"ecommerce/user"
	"fmt"
	"log"
)

func ServeGo() {
	cfg, errs := config.Load()
	if errs != nil {
		panic(errs)
	}

	fmt.Println(cfg.Port)
	fmt.Println(cfg.JWTSecret)

	dbConn, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer dbConn.Close()

	if err := db.RunMigrations(cfg); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	userRepo := user.NewRepository(dbConn)
	userService := user.NewService(userRepo)
	userHandler := userctrl.NewHandler(userService)

	productRepo := product.NewProductRepository(dbConn)
	productService := product.NewService(productRepo)
	productHandler := pdCtrl.NewHandler(productService)

	rest.NewServer(userHandler, productHandler).Start(cfg)
}
