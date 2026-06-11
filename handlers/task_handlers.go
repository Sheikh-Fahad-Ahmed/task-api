package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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

func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Printf("Cannot convert string id to int: %s", idString)
	}

	w.Header().Set("Content-Type", "application/json")
	tasks := store.GetAll()

	if len(tasks) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}

	for _, task := range tasks {
		if task.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"error": "Task ID not found"})
}
