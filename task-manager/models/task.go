package models

import "time"

// TaskStatus represents the status of a task
type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCompleted  TaskStatus = "completed"
)

// TaskPriority represents the priority of a task
type TaskPriority string

const (
	TaskPriorityLow    TaskPriority = "low"
	TaskPriorityMedium TaskPriority = "medium"
	TaskPriorityHigh   TaskPriority = "high"
)

// Task represents a task entity
type Task struct {
	ID          string       `bson:"_id,omitempty" json:"id"`
	Title       string       `bson:"title" json:"title"`
	Description string       `bson:"description" json:"description"`
	UserID      string       `bson:"user_id" json:"user_id"`
	Status      string   `bson:"status" json:"status"`
	// Priority    TaskPriority `bson:"priority" json:"priority"`
	DueDate     Date         `bson:"due_date" json:"due_date"`
	CreatedAt   time.Time    `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time    `bson:"updated_at" json:"updated_at"`
} 