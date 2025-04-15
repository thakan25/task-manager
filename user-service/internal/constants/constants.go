package constants

const (
	// Database constants
	UserDBName     = "userdb"
	UserCollection = "users"

	// API routes
	BasePath      = "/users"
	EmailPath     = "/email/:email"
	IDPath        = "/:id"

	// HTTP methods
	POST = "POST"
	GET  = "GET"

	// Environment variables
	MongoDBURI = "MONGODB_URI"
	Port       = "PORT"

	// Default values
	DefaultPort = "8080"

	// Error messages
	ErrEmailRequired    = "email is required"
	ErrUserIDRequired   = "user id is required"
	ErrUserNotFound     = "user not found"
	ErrUserExists       = "user already exists"
	ErrInvalidRequest   = "invalid request"
	ErrInternalServer   = "internal server error"

	// Success messages
	MsgUserCreated = "user created successfully"
) 