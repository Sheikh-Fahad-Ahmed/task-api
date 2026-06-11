package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Sheikh-Fahad-Ahmed/task-api/models"
	"github.com/Sheikh-Fahad-Ahmed/task-api/store"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tasks := store.GetAll()

	json.NewEncoder(w).Encode(tasks)
	w.WriteHeader(http.StatusOK)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newTaskInput models.TaskInput

	err := json.NewDecoder(r.Body).Decode(&newTaskInput)
	if err != nil {
		log.Fatal("Cannot decode response body")
	}

	tasks := store.GetAll()

	var id int
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	} else {
		id = 1
	}

	newTask := models.New(id, newTaskInput.Title, newTaskInput.Description, newTaskInput.Status)
	newTask.CreatedAt = time.Now()

	store.Add(*newTask)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(*newTask)
}
