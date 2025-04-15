package utils

import (
	"errors"
	"strings"

	"user-service/internal/constants"
)

// ValidateEmail checks if the email is valid
func ValidateEmail(email string) error {
	if email == "" {
		return errors.New(constants.ErrEmailRequired)
	}
	if !strings.Contains(email, "@") {
		return errors.New("invalid email format")
	}
	return nil
}

// ValidateUserID checks if the user ID is valid
func ValidateUserID(id string) error {
	if id == "" {
		return errors.New(constants.ErrUserIDRequired)
	}
	return nil
}

// ValidateCreateUserRequest validates the create user request
func ValidateCreateUserRequest(email string) error {
	if err := ValidateEmail(email); err != nil {
		return err
	}
	return nil
} 