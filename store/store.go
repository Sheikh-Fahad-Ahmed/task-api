package store

import (
	"time"

	"github.com/Sheikh-Fahad-Ahmed/task-api/models"
)

var tasks []models.Task

var nextID = 1

func GetAll() []models.Task {
	if tasks != nil {
		return tasks
	}
	return []models.Task{}
}

func Add(newTaskInput models.TaskInput) models.Task {
	var id int
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	} else {
		id = 1
	}

	newTask := models.New(id, newTaskInput.Title, newTaskInput.Description, newTaskInput.Status)
	newTask.CreatedAt = time.Now()

	tasks = append(tasks, *newTask)
	return *newTask
}
