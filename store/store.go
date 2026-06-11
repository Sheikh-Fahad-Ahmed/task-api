package store

import "github.com/Sheikh-Fahad-Ahmed/task-api/models"

var tasks []models.Task

var nextID = 1

func GetAll() []models.Task {
	if tasks != nil {
		return tasks
	}
	return []models.Task{}
}

func Add(newTask models.Task) {
	tasks = append(tasks, newTask)
}
