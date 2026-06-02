package main

import (
	"ecommerce/servego"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	fmt.Println(os.Getenv("PORT"))
	fmt.Println(os.Getenv("DB_NAME"))
	servego.ServeGo()
}
