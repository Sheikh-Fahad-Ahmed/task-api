package handlers

import (
	"encoding/json"

	"net/http"

	"github.com/Sheikh-Fahad-Ahmed/task-api/models"
	"github.com/Sheikh-Fahad-Ahmed/task-api/store"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func GetTasks(c *gin.Context) {
	tasks := store.GetAll()
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var newTaskInput models.TaskInput

	err := c.ShouldBindJSON(&newTaskInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask, err := store.Add(newTaskInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

func GetTaskByID(c *gin.Context) {
	idString := c.Param("id")

	task, err := store.GetByID(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	idString := c.Param("id")

	var TaskInput models.TaskInput
	err := c.ShouldBindJSON(&TaskInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	task, err := store.Update(idString, TaskInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, task)

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	w.Header().Set("Content-Type", "application/json")

	err := store.Delete(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "task deleted successfully"})

}
