package api

import (
	"fmt"
	"log"
	"os"

	"github.com/agusdwis/blog-app/api/controllers"
	"github.com/agusdwis/blog-app/api/seed"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error load godotenv: %v", err)
	} else {
		fmt.Println("Success get env")
	}

	server.Initialize(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(server.DB)

	server.Run(":8080")
}

