package config

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config holds all configuration for the application
type Config struct {
	MongoDB struct {
		URI      string
		Database string
		Timeout  time.Duration
	}
	Server struct {
		Port           string
		LogLevel       string
		EnableCORS     bool
		MaxRequestSize int64
	}
}

// LoadConfig loads configuration from environment variables with defaults
func LoadConfig() *Config {
	return &Config{
		MongoDB: struct {
			URI      string
			Database string
			Timeout  time.Duration
		}{
			URI:      getEnvString("MONGODB_URI", "mongodb://localhost:27017"),
			Database: getEnvString("MONGODB_DATABASE", "task_manager"),
			Timeout:  time.Duration(getEnvInt("MONGODB_TIMEOUT", 10)) * time.Second,
		},
		Server: struct {
			Port           string
			LogLevel       string
			EnableCORS     bool
			MaxRequestSize int64
		}{
			Port:           getEnvString("PORT", "8080"),
			LogLevel:       getEnvString("LOG_LEVEL", "info"),
			EnableCORS:     getEnvBool("ENABLE_CORS", true),
			MaxRequestSize: getEnvInt64("MAX_REQUEST_SIZE", 1048576), // 1MB default
		},
	}
}

// getEnvString gets a string value from environment variable
func getEnvString(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvInt gets an integer value from environment variable
func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
		log.Printf("Error converting %s to int, using default value", key)
	}
	return defaultValue
}

// getEnvBool gets a boolean value from environment variable
func getEnvBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
		log.Printf("Error converting %s to bool, using default value", key)
	}
	return defaultValue
}

// getEnvInt64 gets an int64 value from environment variable
func getEnvInt64(key string, defaultValue int64) int64 {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// ConnectMongoDB connects to MongoDB using the provided configuration
func ConnectMongoDB(cfg *Config) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.MongoDB.URI))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.MongoDB.Timeout)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to MongoDB")
	return client, nil
} 