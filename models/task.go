package models

import "time"

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Status      string `json:"status" binding:"required"`
}

func New(id int, title string, description string, status string) *Task {
	return &Task{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
	}
}
