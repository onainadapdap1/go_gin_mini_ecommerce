package main

import (
	"fmt"
	"go_gin_mini_ecommerce/routers"
	"log"

	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load();err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	fmt.Println("Main application starts")
	loadEnv()

	log.Fatal(routers.RunAPI(":8080"))
}