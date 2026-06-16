package servego

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/controller/product"
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

	productRepository := repo.NewProductRepo(dbConn)
	productHandler := product.NewHandler(productRepository)

	rest.NewServer(userHandler, productHandler).Start(cfg)
}
