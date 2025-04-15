package mongodb

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/SachinThakan/task-manager/models"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TaskRepository implements the TaskRepository interface for MongoDB
type TaskRepository struct {
	collection *mongo.Collection
}

// NewTaskRepository creates a new MongoDB task repository
func NewTaskRepository(client *mongo.Client, dbName string) *TaskRepository {
	return &TaskRepository{
		collection: client.Database(dbName).Collection("tasks"),
	}
}

// Create creates a new task in the database
func (r *TaskRepository) Create(ctx context.Context, task *models.Task) error {
	// Set timestamps
	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	// Insert task
	_, err := r.collection.InsertOne(ctx, task)
	return err
}

// GetByID retrieves a task by its ID
func (r *TaskRepository) GetByID(ctx context.Context, taskId string) (*models.Task, error) {
	// Find task
	var task models.Task
	err := r.collection.FindOne(ctx, bson.M{"_id": taskId}).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

// GetAll retrieves all tasks
func (r *TaskRepository) GetAll(ctx context.Context) ([]*models.Task, error) {
	// Find all tasks
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode tasks
	var tasks []*models.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// Update updates a task
func (r *TaskRepository) Update(ctx context.Context, task *models.Task) error {
	// Set updated timestamp
	task.UpdatedAt = time.Now()

	// Update task
	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": task.ID}, task)
	return err
}

// Delete deletes a task
func (r *TaskRepository) Delete(ctx context.Context, taskId string) error {
	// Delete task
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": taskId})
	return err			
}


func (r *TaskRepository) GetByStatus(ctx context.Context, status models.TaskStatus) ([]*models.Task, error) {
	filter := bson.M{"status": string(status)}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []*models.Task
	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, cursor.Err()
}

func (r *TaskRepository) GetPaginatedTasks(ctx context.Context, userID string, status models.TaskStatus, page, limit int) ([]*models.Task, error) {
	filter := bson.M{}
	if userID != "" {
		filter["user_id"] = userID
	}
	if status != "" {
		filter["status"] = status
	}

	skip := int64((page - 1) * limit)
	limit64 := int64(limit)

	findOptions := options.Find().SetSkip(skip).SetLimit(limit64)

	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []*models.Task
	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, cursor.Err()
}