package main

import (
	"net/http"

	"github.com/agustinalonsocantoli/go-api/database"
	"github.com/agustinalonsocantoli/go-api/models"
	"github.com/agustinalonsocantoli/go-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	database.DatabaseConnection()

	database.DB.AutoMigrate(models.Task{})
	database.DB.AutoMigrate(models.User{})

	router := mux.NewRouter()

	router.HandleFunc("/users", routes.UsersIndex).Methods("GET")
	router.HandleFunc("/users/{id}", routes.UsersShow).Methods("GET")
	router.HandleFunc("/users", routes.UsersStore).Methods("POST")
	router.HandleFunc("/users/{id}", routes.UsersDelete).Methods("DELETE")

	router.HandleFunc("/tasks", routes.TasksIndex).Methods("GET")
	router.HandleFunc("/tasks/{id}", routes.TasksShow).Methods("GET")
	router.HandleFunc("/tasks", routes.TasksStore).Methods("POST")
	router.HandleFunc("/tasks/{id}", routes.TasksDelete).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
