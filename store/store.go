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

func (s *Store) Add(newTaskInput models.TaskInput) (models.Task, error) {
	var id int

	if newTaskInput.Title == "" {
		return models.Task{}, errors.New("Title is required")
	}

	tasks, err := s.GetAll()
	if err != nil {
		return models.Task{}, err
	}

	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	} else {
		id = 1
	}

	title := strings.TrimSpace(newTaskInput.Title)
	newTask := models.New(id, title, newTaskInput.Description, newTaskInput.Status)
	newTask.CreatedAt = time.Now()

	_, err = s.db.Exec("INSERT INTO tasks (id, title, description, status, created_at) VALUES (?, ?, ?, ?, ?) ",
		newTask.ID, newTask.Title, newTask.Description, newTask.Status, newTask.CreatedAt)

	return *newTask, err
}

func (s *Store) GetByID(idString string) (models.Task, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return models.Task{}, errors.New("unable to convert id from string -> int ")
	}

	var task models.Task
	row := s.db.QueryRow("SELECT * FROM tasks WHERE id = ?", id)
	err = row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt)
	return task, err
}

func (s *Store) Update(idString string, taskInput models.TaskInput) (models.Task, error) {
	task, err := s.GetByID(idString)
	if taskInput.Title != "" {
		task.Title = strings.TrimSpace(taskInput.Title)
	}

	if taskInput.Description != "" {
		task.Description = taskInput.Description
	}
	if taskInput.Status != "" {
		task.Status = taskInput.Status
	}

	task.CreatedAt = time.Now()
	id, err := strconv.Atoi(idString)
	if err != nil {
		return models.Task{}, errors.New("unable to convert id from string -> int ")
	}

	_, err = s.db.Exec("UPDATE tasks SET title = ?, description = ?, status = ?, created_at = ? WHERE id = ?",
		task.Title, task.Description, task.Status, task.CreatedAt, id)
	return task, err
}

func (s *Store) Delete(idString string) error {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return errors.New("unable to convert id string -> int")
	}

	row, err := s.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 0 {
		return nil
	}

	return errors.New("Task not found")
}
