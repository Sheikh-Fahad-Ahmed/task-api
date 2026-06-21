package handlers

import (
	"errors"
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

func (h *Handler) GetTasks(c *gin.Context) {
	tasks, err := h.taskStore.GetAll()
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) CreateTask(c *gin.Context) {
	var newTaskInput models.TaskInput

	err := c.ShouldBindJSON(&newTaskInput)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask, err := h.taskStore.Add(newTaskInput)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

func (h *Handler) GetTaskByID(c *gin.Context) {
	idString := c.Param("id")

	task, err := h.taskStore.GetByID(idString)
	if err != nil {
		c.Error(errors.New("ID does not exist."))
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID does not exist."})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) UpdateTask(c *gin.Context) {
	idString := c.Param("id")

	var taskUpdate models.TaskUpdate
	err := c.ShouldBindJSON(&taskUpdate)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.taskStore.Update(idString, taskUpdate)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *Handler) DeleteTask(c *gin.Context) {
	idString := c.Param("id")

	err := h.taskStore.Delete(idString)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted Successfully"})
}
