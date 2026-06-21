package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Sheikh-Fahad-Ahmed/task-api/db"
	"github.com/Sheikh-Fahad-Ahmed/task-api/handlers"
	"github.com/Sheikh-Fahad-Ahmed/task-api/logger"
	"github.com/Sheikh-Fahad-Ahmed/task-api/store"
	"github.com/gin-gonic/gin"
)

func main() {
	logFile := "gin.log"

	database, err := db.InitDB("./tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	taskStore := store.New(database)
	taskHandler := handlers.New(taskStore)

	gin.DisableConsoleColor()
	f, _ := os.Create(logFile)
	gin.DefaultWriter = io.MultiWriter(f)
	log.SetOutput(gin.DefaultWriter)

	r := gin.New()
	r.Use(logger.CustomLogger())
	r.Use(gin.Recovery())

	r.GET("/tasks", taskHandler.GetTasks)
	r.GET("/tasks/:id", taskHandler.GetTaskByID)
	r.POST("/tasks", taskHandler.CreateTask)
	r.PUT("/tasks/:id", taskHandler.UpdateTask)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)

	fmt.Println("server starting on http://localhost:8080")
	fmt.Printf("\nAll logs are routed to: %s", logFile)

	r.Run(":8080")

}
