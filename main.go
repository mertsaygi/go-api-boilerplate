package main

import (
	"fmt"
	"github.com/mertsaygi/go-api-boilerplate/src/config"
	"github.com/mertsaygi/go-api-boilerplate/src/controller"
	"github.com/mertsaygi/go-api-boilerplate/src/models"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	var db *gorm.DB = config.ConnectDB()
	db.AutoMigrate(&models.User{})

	defer config.DisconnectDB(db)

	router := mux.NewRouter()

	router.HandleFunc("/health", controller.HealthCheckHandler)

	router.HandleFunc("/users", controller.ListUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", controller.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", controller.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", controller.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", controller.DeleteUserHandler).Methods("DELETE")

	port := os.Getenv("PORT")
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
