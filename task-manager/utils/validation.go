package utils

import (
	"regexp"
	"errors"
)

var (
	ErrInvalidEmail     = errors.New("invalid email format")
	ErrPasswordTooShort = errors.New("password must be at least 6 characters")
	ErrUsernameTooShort = errors.New("username must be at least 3 characters")
	ErrInvalidRequest   = errors.New("invalid request")
)

// EmailRegex is the regular expression for validating email addresses
var EmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// ValidateEmail validates an email address
func ValidateEmail(email string) error {
	if !EmailRegex.MatchString(email) {
		return ErrInvalidEmail
	}
	return nil
}

// ValidatePassword validates a password
func ValidatePassword(password string) error {
	if len(password) < 6 {
		return ErrPasswordTooShort
	}
	return nil
}

// ValidateUsername validates a username
func ValidateUsername(username string) error {
	if len(username) < 3 {
		return ErrUsernameTooShort
	}
	return nil
}

func ValidateEmailError(email string) error {
	if email == "" {
		return ErrInvalidRequest
	}
	
	// Basic email regex validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return ErrInvalidRequest
	}
	
	return nil
}

func ValidatePasswordError(password string) error {
	if password == "" {
		return ErrInvalidRequest
	}
	
	// Password should be at least 6 characters
	if len(password) < 6 {
		return ErrPasswordTooShort
	}
	
	return nil
}

func ValidateUsernameError(username string) error {
	if username == "" {
		return ErrInvalidRequest
	}
	
	// Username should be at least 3 characters
	if len(username) < 3 {
		return ErrUsernameTooShort
	}
	
	return nil
} 