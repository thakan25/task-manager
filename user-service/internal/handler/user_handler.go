package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"user-service/internal/adapters"
	"user-service/internal/constants"
	"user-service/internal/models/request"
	"user-service/internal/service"
)

// UserHandler handles HTTP requests for user operations
type UserHandler struct {
	userService service.UserService
	adapter     *adapters.UserAdapter
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		adapter:     adapters.NewUserAdapter(),
	}
}

// CreateUser handles the creation of a new user
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req request.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, h.adapter.ToErrorResponse(err))
		return
	}

	dto := h.adapter.ToCreateUserDTO(&req)
	response, err := h.userService.CreateUser(c.Request.Context(), dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, h.adapter.ToErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, h.adapter.ToSuccessResponse(constants.MsgUserCreated, h.adapter.ToUserResponse(response)))
}

// GetUser handles retrieving a user by email
func (h *UserHandler) GetUser(c *gin.Context) {
	var req request.GetUserByEmailRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, h.adapter.ToErrorResponse(err))
		return
	}

	response, err := h.userService.GetUser(c.Request.Context(), req.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, h.adapter.ToErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, h.adapter.ToSuccessResponse("User found", h.adapter.ToUserResponse(response)))
}

// GetUserByID handles retrieving a user by ID
func (h *UserHandler) GetUserByID(c *gin.Context) {
	var req request.GetUserByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, h.adapter.ToErrorResponse(err))
		return
	}

	response, err := h.userService.GetUserByID(c.Request.Context(), req.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, h.adapter.ToErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, h.adapter.ToSuccessResponse("User found", h.adapter.ToUserResponse(response)))
}