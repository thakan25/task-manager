package dtos

import (
	"testing"
	"time"
)

func TestUserDTO(t *testing.T) {
	// Test creating a UserDTO
	user := &UserDTO{
		ID:        "test-id",
		Username:  "test-username",
		Email:     "test@example.com",
		Password:  "test-password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if user.ID != "test-id" {
		t.Errorf("Expected ID to be 'test-id', got '%s'", user.ID)
	}
	if user.Username != "test-username" {
		t.Errorf("Expected Username to be 'test-username', got '%s'", user.Username)
	}
	if user.Email != "test@example.com" {
		t.Errorf("Expected Email to be 'test@example.com', got '%s'", user.Email)
	}
	if user.Password != "test-password" {
		t.Errorf("Expected Password to be 'test-password', got '%s'", user.Password)
	}
} 