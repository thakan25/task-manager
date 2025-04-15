package dtos

import (
	"time"

	"github.com/SachinThakan/task-manager/models"
)

// TaskDTO represents the task data transfer object for service layer
type TaskDTO struct {
	ID          string     `json:"id"`
	UserID      string     `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	DueDate     models.Date `json:"due_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// CreateTaskDTO represents the DTO for creating a task in service layer
type CreateTaskDTO struct {
	UserID      string     `json:"user_id" binding:"required"`
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description"`
	DueDate     models.Date `json:"due_date" binding:"required"`
}

// UpdateTaskDTO represents the DTO for updating a task in service layer
type UpdateTaskDTO struct {
	ID          string     `json:"id" binding:"required"`
	UserID      string     `json:"user_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	DueDate     models.Date `json:"due_date"`
} 