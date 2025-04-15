package errors

import "errors"

var (
	// Common errors
	ErrNotFound = errors.New("not found")

	// User related errors
	ErrInvalidUserEmail = errors.New("invalid email format")
	ErrUserNotFound     = errors.New("user not found")

	// Task related errors
	ErrInvalidTaskStatus = errors.New("invalid task status")
	ErrInvalidTaskTitle  = errors.New("invalid task title")
	ErrInvalidDueDate    = errors.New("invalid due date")
	ErrInvalidUserID     = errors.New("invalid user id")
) 