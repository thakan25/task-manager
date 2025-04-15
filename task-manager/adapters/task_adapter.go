package adapters

import (
	"github.com/SachinThakan/task-manager/models"
	"github.com/SachinThakan/task-manager/models/dtos"
	"github.com/SachinThakan/task-manager/models/requests"
	"github.com/SachinThakan/task-manager/models/responses"
)

// ControllerToServiceAdapter handles conversion between controller requests and service entities
type ControllerToServiceAdapter struct{}

// NewControllerToServiceAdapter creates a new controller to service adapter
func NewControllerToServiceAdapter() *ControllerToServiceAdapter {
	return &ControllerToServiceAdapter{}
}

// toTaskDto converts a create task request to a task entity
func (a *ControllerToServiceAdapter) ToCreateTaskDTO(req requests.CreateTaskRequest) dtos.CreateTaskDTO {
	return dtos.CreateTaskDTO{
		Title:       req.Title,
		Description: req.Description,
		UserID:      req.UserID,
		// Priority:    models.TaskPriority(req.Priority),
		DueDate:     req.DueDate,
	}
}

// ToTaskEntityFromUpdate converts an update task request to a task entity
func (a *ControllerToServiceAdapter) ToTaskDtoForUpdate(req requests.UpdateTaskRequest) dtos.UpdateTaskDTO {
	return dtos.UpdateTaskDTO{
		Title:       req.Title,
		Description: req.Description,
		UserID: 	 req.UserID,
		DueDate:     req.DueDate,
		Status:      req.Status,
	}
}

// ToTaskResponse converts a task entity to a task response
func (a *ControllerToServiceAdapter) ToTaskResponse(task *dtos.TaskDTO) responses.TaskResponse {
	return responses.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      string(task.Status),
		// Priority:    string(task.Priority),
		DueDate:     task.DueDate.ToTime(),
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

// ToTasksResponse converts a slice of task entities to a slice of task responses
func (a *ControllerToServiceAdapter) ToTasksResponse(tasks []*dtos.TaskDTO) []responses.TaskResponse {
	responses := make([]responses.TaskResponse, len(tasks))
	for i, task := range tasks {
		responses[i] = a.ToTaskResponse(task)
	}
	return responses
}

// ServiceToRepositoryAdapter converts service DTO to repository entity
type ServiceToRepositoryAdapter struct{}

func NewServiceToRepositoryAdapter() *ServiceToRepositoryAdapter {
	return &ServiceToRepositoryAdapter{}
}

func (a *ServiceToRepositoryAdapter) ToTaskEntity(dto dtos.TaskDTO) *models.Task {
	return &models.Task{
		ID:          dto.ID,
		Title:       dto.Title,
		UserID:      dto.UserID,
		Description: dto.Description,
		Status:      dto.Status,
		DueDate:     dto.DueDate,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
	}
}

func (a *ServiceToRepositoryAdapter) ToTaskDTO(entity *models.Task) dtos.TaskDTO {
	return dtos.TaskDTO{
		ID:          entity.ID,
		Title:       entity.Title,
		UserID:      entity.UserID,
		Description: entity.Description,
		Status:      string(entity.Status),
		DueDate:     entity.DueDate,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func (a *ServiceToRepositoryAdapter) ToTaskEntities(dtos []dtos.TaskDTO) []*models.Task {
	entities := make([]*models.Task, len(dtos))
	for i, dto := range dtos {
		entities[i] = a.ToTaskEntity(dto)
	}
	return entities
}

func (a *ServiceToRepositoryAdapter) ToTaskDTOs(entities []*models.Task) []dtos.TaskDTO {
	dtos := make([]dtos.TaskDTO, len(entities))
	for i, entity := range entities {
		dtos[i] = a.ToTaskDTO(entity)
	}
	return dtos
} 