package main

import (
	"log"

	"github.com/Sheikh-Fahad-Ahmed/task-api/db"
	"github.com/Sheikh-Fahad-Ahmed/task-api/handlers"
	"github.com/Sheikh-Fahad-Ahmed/task-api/store"
	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime)
	database, err := db.InitDB("./tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	taskStore := store.New(database)

	taskHandler := handlers.New(taskStore)

	r := gin.Default()

	r.GET("/tasks", taskHandler.GetTasks)

	r.GET("tasks/:id", taskHandler.GetTaskByID)

	r.POST("/tasks", taskHandler.CreateTask)

	r.PUT("/tasks/:id", taskHandler.UpdateTask)

	r.DELETE("tasks/:id", handlers.DeleteTask)

	// loggedRoute := logger.Logging(r.ServeHTTP)

	r.Run()

}
