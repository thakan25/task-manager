package service

import (
	"context"
	"errors"
	"log"

	"user-service/internal/constants"
	"user-service/internal/models"
	"user-service/internal/repository"
	"user-service/internal/utils"
)

// UserService defines the interface for user business logic
type UserService interface {
	CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.UserResponse, error)
	GetUser(ctx context.Context, email string) (*models.UserResponse, error)
	GetUserByID(ctx context.Context, id string) (*models.UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.UserResponse, error) {
	log.Println("[CreateUser] Received request to create user with email:", req.Email)

	// Validate request
	if err := utils.ValidateCreateUserRequest(req.Email); err != nil {
		log.Println("[CreateUser] Validation failed:", err)
		return nil, err
	}

	// Check if user already exists
	exists, err := s.userRepo.Exists(ctx, req.Email)
	if err != nil {
		log.Println("[CreateUser] Failed to check if user exists:", err)
		return nil, err
	}
	if exists {
		log.Println("[CreateUser] User already exists with email:", req.Email)
		return nil, errors.New(constants.ErrUserExists)
	}

	// Create new user
	user := &models.User{
		ID:        utils.GenerateUserID(),
		Email:     req.Email,
	}

	log.Println("[CreateUser] Creating user with ID:", user.ID)

	err = s.userRepo.Create(ctx, user)
	if err != nil {
		log.Println("[CreateUser] Failed to create user:", err)
		return nil, err
	}

	log.Println("[CreateUser] Successfully created user:", user.ID)

	return &models.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (s *userService) GetUser(ctx context.Context, email string) (*models.UserResponse, error) {
	// Validate email
	if err := utils.ValidateEmail(email); err != nil {
		return nil, err
	}

	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New(constants.ErrUserNotFound)
	}

	return &models.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*models.UserResponse, error) {
	// Validate user ID
	if err := utils.ValidateUserID(id); err != nil {
		return nil, err
	}

	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New(constants.ErrUserNotFound)
	}

	return &models.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}