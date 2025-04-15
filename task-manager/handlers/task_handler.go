package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/SachinThakan/task-manager/adapters"
	// "github.com/SachinThakan/task-manager/models"
	// "github.com/SachinThakan/task-manager/models/dtos"
	"github.com/SachinThakan/task-manager/models/requests"
	"github.com/SachinThakan/task-manager/service"
	
)

// TaskHandler handles HTTP requests for tasks
type TaskHandler struct {
	service *service.TaskService
	adapter *adapters.ControllerToServiceAdapter
}

// NewTaskHandler creates a new task handler
func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
		adapter: adapters.NewControllerToServiceAdapter(),
	}
}

// CreateTask handles the creation of a new task
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req requests.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert request to entity
	task := h.adapter.ToCreateTaskDTO(req)

	// Create task
	taskDto, err := h.service.CreateTask(c.Request.Context(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert entity to response
	response := h.adapter.ToTaskResponse(taskDto)
	c.JSON(http.StatusCreated, response)
}

// GetTask handles retrieving a single task by ID
func (h *TaskHandler) GetTasks(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	status := c.Query("status")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	tasks, err := h.service.GetTasks(c.Request.Context(), userID, status, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tasks"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// UpdateTask handles updating a task
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var req requests.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert request to entity
	task := h.adapter.ToTaskDtoForUpdate(req)
	task.ID = id
	// Update task
	taskDto, err := h.service.UpdateTask(c.Request.Context(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert entity to response
	response := h.adapter.ToTaskResponse(taskDto)
	c.JSON(http.StatusOK, response)
}

// DeleteTask handles deleting a task
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteTask(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
} 