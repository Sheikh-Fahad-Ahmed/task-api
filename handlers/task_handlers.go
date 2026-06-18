package handlers

import (
	"net/http"

	"github.com/Sheikh-Fahad-Ahmed/task-api/models"
	"github.com/Sheikh-Fahad-Ahmed/task-api/store"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	taskStore *store.Store
}

func New(taskStore *store.Store) *Handler {
	return &Handler{taskStore: taskStore}
}

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
		return
	}

	task, err := store.Update(idString, TaskInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)

}

func DeleteTask(c *gin.Context) {
	idString := c.Param("id")

	err := store.Delete(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted Successfully"})

}
