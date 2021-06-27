package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/agusdwis/blog-app/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

type Server struct {
	DB		*gorm.DB
	Router	*mux.Router
}

func (server *Server) Initialize(DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbPort, DbHost, DbName)

	server.DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to database")
		log.Fatal("This is the error", err)
	} else {
		fmt.Printf("We are connected to database")
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{})

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
