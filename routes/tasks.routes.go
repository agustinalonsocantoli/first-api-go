package routes

import (
	"encoding/json"
	"net/http"

	"github.com/agustinalonsocantoli/go-api/database"
	"github.com/agustinalonsocantoli/go-api/models"
	"github.com/gorilla/mux"
)

func TasksIndex(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	findTasks := database.DB.Preload("User").Find(&tasks)

	if err := findTasks.Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&tasks)
}

func TasksShow(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	if err := database.DB.Preload("User").First(&task, params["id"]).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func TasksStore(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	json.NewDecoder(r.Body).Decode(&task)
	createTask := database.DB.Create(&task)

	if err := createTask.Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&task)
}

func TasksDelete(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	if err := database.DB.First(&task, params["id"]).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	if err := database.DB.Delete(&task).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	} else {
		json.NewEncoder(w).Encode(&task)
	}
}
