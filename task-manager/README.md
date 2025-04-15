# task manager Service

A microservice for task management built with Go and MongoDB.

## Features

- User CRUD operations (Create, Read, Update, Delete)
- MongoDB integration for data persistence
- RESTful API with JSON responses
- Environment-based configuration
- Health check endpoint

## Prerequisites

- Go 1.18 or later
- MongoDB 4.4 or later

## Configuration

The service can be configured using environment variables:

```env
PORT=8081                            # Server port (default: 8081)
MONGO_URI=mongodb://localhost:27017  # MongoDB connection URI
DB_NAME=task_manager                # MongoDB database name
```

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd user-service
```

2. Install dependencies:
```bash
go mod download
```

3. Run the service:
```bash
go run main.go
```

## API Endpoints

### Health Check
```
GET /api/v1/health
Response: 200 OK
```


### create task
```
POST /api/v1/tasks

create task with following sample date:

{
  "title": "Finish project documentation",
  "description": "Complete the final project documentation for the client.",
  "user_id": "user123",
  "due_date": "2025-04-20"
}

```


### update task
```
PUT /api/v1/tasks/:id

update task with following sample date:

{
  "title": "Finish project documentation - updated",
  "description": "Update the final project documentation with additional features.",
  "user_id": "user123",
  "status": "in-progress",
  "due_date": "2025-04-25"
}
```


### get tasks(Paginated)
```
GET /api/v1/tasks?user_id=user1&status=pending&page=1&limit=10

fetch paginated tasks for a gven userId and status

room for improvement: filter can be enhanced to sort the records, to support multiple statuses
```

### Delete task
```
DELETE /api/v1/tasks/:id

delete task by taskId
```


## Error Responses

The API returns appropriate HTTP status codes and error messages:

- 400 Bad Request: Invalid input data
- 500 Internal Server Error: Server-side error

## Project Structure

```
.
├── config/         # Configuration management
├── handlers/       # HTTP request handlers
├── models/         # Data models and DTOs
├── repository/     # Data access layer
│   └── mongodb/   # MongoDB implementation
├── service/        # Business logic
└── main.go        # Application entry point
```

## License

MIT
