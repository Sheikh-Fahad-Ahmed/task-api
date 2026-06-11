package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Sheikh-Fahad-Ahmed/task-api/store"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tasks := store.GetAll() 

	json.NewEncoder(w).Encode(tasks)
}

