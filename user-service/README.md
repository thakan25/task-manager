# User Service

A Go microservice for user management with MongoDB integration.

## Features

- Create users with email
- Get user details by email
- Validate user existence
- Clean architecture with separation of concerns
- MongoDB integration for persistence

## Prerequisites

- Go 1.21 or later
- MongoDB running locally or accessible via network
- Make sure MongoDB is running on the default port (27017) or update the MONGODB_URI in .env

## Setup

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Create a .env file with the following variables:
   ```
   MONGODB_URI=mongodb://localhost:27017
   PORT=8080
   ```

## Running the Service

```bash
go run main.go
```

The service will start on port 8080 by default.

## API Endpoints

### Create User
```
POST /users
Content-Type: application/json

{
    "email": "user@example.com"
}
```

### Get User
```
GET /users/:id
```

### Validate User
```
GET /users/validate/:email
```

## Project Structure

```
.
├── internal/
│   ├── handler/     # HTTP handlers
│   ├── models/      # Domain models and DTOs
│   ├── repository/  # MongoDB operations
│   └── service/     # Business logic
├── main.go          # Application entry point
├── go.mod           # Go module file
└── .env             # Environment configuration
```