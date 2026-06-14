package servego

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/controller/product"
	"ecommerce/rest/controller/user"
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

	userRepository := repo.NewUserRepo(dbConn)
	productRepository := repo.NewProductRepo(dbConn)
	loginRepository := repo.NewLoginRepo(dbConn)

	productHandler := product.NewHandler(productRepository)

	userHandler := user.NewHandler(userRepository, loginRepository)

	// reviewHandler := review.NewHandler(reviewRepository)

	rest.NewServer(userHandler, productHandler).Start(cfg)
}
