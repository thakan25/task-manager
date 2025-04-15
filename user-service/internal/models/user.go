package models

import (
	"time"
)

// User represents the domain model for a user
type User struct {
	ID        string    `bson:"_id" json:"id"`
	Email     string    `bson:"email" json:"email"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

// CreateUserRequest represents the DTO for creating a new user
type CreateUserRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// UserResponse represents the DTO for user responses
type UserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// ValidateUserResponse represents the DTO for user validation response
type ValidateUserResponse struct {
	IsValid bool   `json:"is_valid"`
	Message string `json:"message"`
} 