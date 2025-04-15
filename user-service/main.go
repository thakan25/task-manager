package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"user-service/internal/constants"
	"user-service/internal/handler"
	"user-service/internal/repository"
	"user-service/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	collection *mongo.Collection
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// MongoDB connection
	clientOptions := options.Client().ApplyURI(os.Getenv(constants.MongoDBURI))
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(constants.UserDBName).Collection(constants.UserCollection)
}

func main() {
	// Initialize repository
	userRepo := repository.NewUserRepository(collection)

	// Initialize service
	userService := service.NewUserService(userRepo)

	// Initialize handler
	userHandler := handler.NewUserHandler(userService)

	// Setup router
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Service is up!"})
	})

	// User routes
	router.POST(constants.BasePath, userHandler.CreateUser)
	router.GET(constants.BasePath+constants.EmailPath, userHandler.GetUser)
	router.GET(constants.BasePath+constants.IDPath, userHandler.GetUserByID)

	// Start server
	port := os.Getenv(constants.Port)
	if port == "" {
		port = constants.DefaultPort
	}
	router.Run(":" + port)
} 