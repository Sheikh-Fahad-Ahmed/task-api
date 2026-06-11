package handlers

import (
	"encoding/json"

	"net/http"

	"github.com/Sheikh-Fahad-Ahmed/task-api/models"
	"github.com/Sheikh-Fahad-Ahmed/task-api/store"
	"github.com/gorilla/mux"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tasks := store.GetAll()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newTaskInput models.TaskInput

	err := json.NewDecoder(r.Body).Decode(&newTaskInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "unable to decode Request body"})
		return
	}

	if newTaskInput.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Title is required"})
		return
	}

	newTask := store.Add(newTaskInput)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]

	w.Header().Set("Content-Type", "application/json")

	task, err := store.GetByID(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}
