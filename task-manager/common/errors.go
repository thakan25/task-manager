package common

import "errors"

var (
	// Task errors
	ErrTaskNotFound = errors.New("task not found")
	ErrInvalidTaskStatus = errors.New("invalid task status")

	// User errors
	ErrUserNotFound = errors.New("user not found")
	ErrEmailExists = errors.New("email already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidRequest = errors.New("invalid request")
) 