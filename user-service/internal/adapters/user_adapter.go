package adapters

import (
	"user-service/internal/models"
	"user-service/internal/models/request"
	"user-service/internal/models/response"
)

// UserAdapter handles conversions between request/response models and DTOs
type UserAdapter struct{}

// NewUserAdapter creates a new instance of UserAdapter
func NewUserAdapter() *UserAdapter {
	return &UserAdapter{}
}

// ToCreateUserDTO converts a CreateUserRequest to CreateUserRequest DTO
func (a *UserAdapter) ToCreateUserDTO(req *request.CreateUserRequest) *models.CreateUserRequest {
	return &models.CreateUserRequest{
		Email: req.Email,
	}
}

// ToUserResponse converts a UserResponse DTO to UserResponse
func (a *UserAdapter) ToUserResponse(dto *models.UserResponse) *response.UserResponse {
	return &response.UserResponse{
		ID:        dto.ID,
		Email:     dto.Email,
		CreatedAt: dto.CreatedAt,
	}
}

// ToErrorResponse creates an ErrorResponse
func (a *UserAdapter) ToErrorResponse(err error) *response.ErrorResponse {
	return &response.ErrorResponse{
		Error: err.Error(),
	}
}

// ToSuccessResponse creates a SuccessResponse with data
func (a *UserAdapter) ToSuccessResponse(message string, data interface{}) *response.SuccessResponse {
	return &response.SuccessResponse{
		Message: message,
		Data:    data,
	}
} 