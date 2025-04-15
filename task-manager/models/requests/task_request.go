package requests

import (
	"github.com/SachinThakan/task-manager/models"
)

// CreateTaskRequest represents a request to create a new task
type CreateTaskRequest struct {
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description" binding:"required"`
	UserID      string     `json:"user_id" binding:"required"`
	Priority    string     `json:"priority" binding:"required"`
	DueDate     models.Date `json:"due_date" binding:"required"`
}

// UpdateTaskRequest represents a request to update an existing task
type UpdateTaskRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	UserID      string     `json:"user_id"`
	Status      string     `json:"status"`
	Priority    string     `json:"priority"`
	DueDate     models.Date `json:"due_date"`
} 