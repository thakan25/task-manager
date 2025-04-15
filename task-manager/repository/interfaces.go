package repository

import (
	"context"
	"github.com/SachinThakan/task-manager/models"
)

// TaskRepository defines the interface for task data access
type TaskRepository interface {
	Create(ctx context.Context, task *models.Task) error
	GetByID(ctx context.Context, id string) (*models.Task, error)
	GetAll(ctx context.Context) ([]*models.Task, error)
	GetByStatus(ctx context.Context, status models.TaskStatus) ([]*models.Task, error)
	Update(ctx context.Context, task *models.Task) error
	Delete(ctx context.Context, id string) error
	GetPaginatedTasks(ctx context.Context, userID string, status models.TaskStatus, page, limit int) ([]*models.Task, error)
}

// UserValidator defines the interface for user validation
type UserValidator interface {
	ValidateUser(ctx context.Context, userID string) (bool, error)
} 