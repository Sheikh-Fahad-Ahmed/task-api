package store

import (
	"errors"
	"strconv"
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

func GetByID(idString string) (models.Task, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return models.Task{}, errors.New("unable to convert id from string -> int ")
	}

	if len(tasks) == 0 {
		return models.Task{}, errors.New("Tasks is empty")
	}

	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return models.Task{}, errors.New("Task does not exists")
}
