package store

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/Sheikh-Fahad-Ahmed/task-api/models"
)

type Store struct {
	db *sql.DB
}

func New(database *sql.DB) *Store {
	return &Store{
		db: database,
	}
}



func (s *Store) GetAll() ([]models.Task, error) {
	rows, err := s.db.Query("SELECT * FROM tasks;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt)
		tasks = append(tasks, t)
	}


	return tasks, rows.Err()
}

func Add(newTaskInput models.TaskInput) (models.Task, error) {
	var id int
	if newTaskInput.Title == "" {
		return models.Task{}, errors.New("Title is required")
	}
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	} else {
		id = 1
	}

	title := strings.TrimSpace(newTaskInput.Title)
	newTask := models.New(id, title, newTaskInput.Description, newTaskInput.Status)
	newTask.CreatedAt = time.Now()

	tasks = append(tasks, *newTask)
	return *newTask, nil
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

func Update(idString string, taskInput models.TaskInput) (models.Task, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return models.Task{}, errors.New("unable to convert id from string -> int")
	}

	var title, description, status string

	for i, task := range tasks {
		if task.ID == id {
			if taskInput.Title != "" {
				title = strings.TrimSpace(taskInput.Title)
			} else {
				title = task.Title
			}
			if taskInput.Description != "" {
				description = taskInput.Description
			} else {
				description = task.Description
			}
			if taskInput.Status != "" {
				status = taskInput.Status
			} else {
				status = task.Status
			}
			tasks[i] = *models.New(id, title, description, status)
			tasks[i].CreatedAt = time.Now()
			return tasks[i], nil
		}
	}

	return models.Task{}, errors.New("Task not Found..")
}

func Delete(idString string) error {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return errors.New("unable to convert id string -> int")
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("Task not found")
}
