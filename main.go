package main

import (
	"log"

	"github.com/Sheikh-Fahad-Ahmed/task-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime)

	r := gin.Default()

	r.GET("/tasks", handlers.GetTasks)
	// r.HandleFunc("/tasks/{id}", handlers.GetTaskByID).Methods("GET")
	r.GET("tasks/:id", handlers.GetTaskByID)

	r.POST("/tasks", handlers.CreateTask)

	// r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	r.PUT("/tasks/:id", handlers.UpdateTask)

	// r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	// loggedRoute := logger.Logging(r.ServeHTTP)

	r.Run()

}

/*
Middleware: add a custom logging middleware (wrap http.Handler to log method, path, and duration) and explain in your README why logging matters
Input sanitization: trim whitespace from title before storing it
Query filter: GET /tasks?status=done filters by status on the list endpoint

*/
