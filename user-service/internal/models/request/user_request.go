package request

// CreateUserRequest represents the request for creating a user
type CreateUserRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// GetUserByEmailRequest represents the request for getting a user by email
type GetUserByEmailRequest struct {
	Email string `uri:"email" binding:"required"`
}

// GetUserByIDRequest represents the request for getting a user by ID
type GetUserByIDRequest struct {
	ID string `uri:"id" binding:"required"`
} 