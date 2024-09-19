package routes

import (
	"encoding/json"
	"net/http"

	"github.com/agustinalonsocantoli/go-api/database"
	"github.com/agustinalonsocantoli/go-api/models"
	"github.com/gorilla/mux"
)

func UsersIndex(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)
}

func UsersShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User

	if err := database.DB.Preload("Tasks").First(&user, params["id"]).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func UsersStore(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)
	createUser := database.DB.Create(&user)
	err := createUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func UsersDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var users models.User

	if err := database.DB.First(&users, params["id"]).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	if errD := database.DB.Delete(&users).Error; errD != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errD.Error()))
		return
	} else {
		json.NewEncoder(w).Encode(&users)
	}

}
