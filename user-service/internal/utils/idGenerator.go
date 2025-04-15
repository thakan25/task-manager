package utils

import (
	"strings"
	"github.com/google/uuid"
)

// GenerateUserID generates a UUID-based user ID with "U" prefix and no hyphens.
func GenerateUserID() string {
	rawUUID := uuid.New().String()
	cleanUUID := strings.ReplaceAll(rawUUID, "-", "")
	return "U" + cleanUUID
}
