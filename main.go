package main

import (
	"log"
	"net/http"

	"github.com/Sheikh-Fahad-Ahmed/task-api/handlers"
	"github.com/Sheikh-Fahad-Ahmed/task-api/logger"
	"github.com/gorilla/mux"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime )

	r := mux.NewRouter()

	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.GetTaskByID).Methods("GET")

	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")

	r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")

	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	loggedRoute := logger.Logging(r.ServeHTTP)

	http.ListenAndServe(":8080", loggedRoute)

}

/*
Middleware: add a custom logging middleware (wrap http.Handler to log method, path, and duration) and explain in your README why logging matters
Input sanitization: trim whitespace from title before storing it
Query filter: GET /tasks?status=done filters by status on the list endpoint

*/
