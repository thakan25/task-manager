package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/SachinThakan/task-manager/config"
	"github.com/SachinThakan/task-manager/constants"
	"github.com/SachinThakan/task-manager/handlers"
	"github.com/SachinThakan/task-manager/accessor"
	"github.com/SachinThakan/task-manager/repository/mongodb"
	"github.com/SachinThakan/task-manager/service"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), cfg.MongoDB.Timeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDB.URI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	// Ping MongoDB
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	log.Println("Connected to MongoDB")

	// Initialize repositories
	taskRepo := mongodb.NewTaskRepository(client, cfg.MongoDB.Database)
	userAccessor := accessor.NewUserServiceAccessor()

	// Initialize services
	taskService := service.NewTaskService(taskRepo, *userAccessor)

	// Initialize handlers
	taskHandler := handlers.NewTaskHandler(taskService)

	// Initialize router
	router := gin.Default()

	// Health check
	router.GET(constants.Health, func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// API routes
	api := router.Group(constants.ApiPath)
	tasks := api.Group(constants.BasePath)
	{
		tasks.POST(constants.DefaultPath, taskHandler.CreateTask)
		tasks.GET(constants.DefaultPath, taskHandler.GetTasks)
		tasks.PUT(constants.IDPath, taskHandler.UpdateTask)
		tasks.DELETE(constants.IDPath, taskHandler.DeleteTask)
	}

	// Start server
	port := os.Getenv(constants.DefaultPort)
	if port == "" {
		port = constants.DefaultPort
	}
	router.Run(":" + port)
} 