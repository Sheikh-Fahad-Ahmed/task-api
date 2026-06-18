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

	r.GET("tasks/:id", handlers.GetTaskByID)

	r.POST("/tasks", handlers.CreateTask)


	r.PUT("/tasks/:id", handlers.UpdateTask)


	r.DELETE("tasks/:id", handlers.DeleteTask)

	// loggedRoute := logger.Logging(r.ServeHTTP)

	r.Run()

}

