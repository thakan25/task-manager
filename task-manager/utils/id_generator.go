package utils

import (
	"strings"

	"github.com/google/uuid"
)

// IDPrefix represents the prefix for different types of IDs
type IDPrefix string

const (
	// TaskIDPrefix is the prefix for task IDs
	TaskIDPrefix IDPrefix = "T"
	// UserIDPrefix is the prefix for user IDs
	UserIDPrefix IDPrefix = "U"
)

// IDGenerator provides methods for generating formatted IDs
type IDGenerator struct{}

// NewIDGenerator creates a new IDGenerator instance
func NewIDGenerator() *IDGenerator {
	return &IDGenerator{}
}

// GenerateID creates a new formatted ID with the given prefix
func (g *IDGenerator) GenerateID(prefix IDPrefix) string {
	uuidStr := uuid.New().String()
	return string(prefix) + strings.ReplaceAll(uuidStr, "-", "")
}

// GenerateTaskID generates a new task ID
func (g *IDGenerator) GenerateTaskID() string {
	return g.GenerateID(TaskIDPrefix)
}

// GenerateUserID generates a new user ID
func (g *IDGenerator) GenerateUserID() string {
	return g.GenerateID(UserIDPrefix)
} 