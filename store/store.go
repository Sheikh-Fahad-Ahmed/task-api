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
	db     *sql.DB
	status map[string]struct{}
}

func New(database *sql.DB) *Store {
	return &Store{
		db: database,
		status: map[string]struct{}{
			"pending":     {},
			"in-progress": {},
			"done":        {},
		},
	}
}

func (s *Store) checkStatus(newTaskInput *models.TaskInput) bool {
	if newTaskInput.Status == "" {
		newTaskInput.Status = "pending"
	}
	_, exists := s.status[newTaskInput.Status]
	return exists

}

func (s *Store) GetAll(status string) ([]models.Task, error) {
	var rows *sql.Rows
	var err error

	if status == "" {
		rows, err = s.db.Query("SELECT id, title, description, status, created_at FROM tasks;")
		if err != nil {
			return nil, err
		}
	} else {
		rows, err = s.db.Query("SELECT id, title, description, status, created_at FROM tasks WHERE status = ?", status)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	tasks := []models.Task{}
	for rows.Next() {
		var t models.Task
		rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt)
		tasks = append(tasks, t)
	}

	return tasks, rows.Err()
}

func (s *Store) Add(newTaskInput models.TaskInput) (models.Task, error) {

	if newTaskInput.Title == "" {
		return models.Task{}, errors.New("Title is required")
	}

	if !s.checkStatus(&newTaskInput) {
		return models.Task{}, errors.New("Invalid status")
	}

	title := strings.TrimSpace(newTaskInput.Title)
	newTask := models.New(title, newTaskInput.Description, newTaskInput.Status)

	result, err := s.db.Exec("INSERT INTO tasks (title, description, status, created_at) VALUES (?, ?, ?, ?) ",
		newTask.Title, newTask.Description, newTask.Status, newTask.CreatedAt)

	id, err := result.LastInsertId()
	if err != nil {
		return models.Task{}, err
	}

	newTask.ID = int(id)
	return *newTask, err
}

func (s *Store) GetByID(idString string) (models.Task, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return models.Task{}, errors.New("unable to convert id from string -> int ")
	}

	var task models.Task
	row := s.db.QueryRow("SELECT id, title, description, status, created_at FROM tasks WHERE id = ?", id)
	err = row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt)
	return task, err
}

func (s *Store) Update(idString string, taskUpdate models.TaskUpdate) (models.Task, error) {
	task, err := s.GetByID(idString)
	if taskUpdate.Title != "" {
		task.Title = strings.TrimSpace(taskUpdate.Title)
	}

	if taskUpdate.Description != "" {
		task.Description = taskUpdate.Description
	}
	if taskUpdate.Status != "" {
		task.Status = taskUpdate.Status
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
