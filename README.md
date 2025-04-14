# Task Manager Application with User Service

This project consists of two microservices:

1. **user-service**: Manages user information.
2. **task-manager**: Manages tasks assigned to users. This service communicates with the user-service to validate user data.

Each service uses its own MongoDB instance and is containerized using Docker.

---

## 🐳 Getting Started

### Prerequisites

- Docker & Docker Compose installed

---

## 🚀 Running the Services

To start both services and their MongoDB instances, run the following two commands from the root directory:

```bash
# Step 1: Build the images
docker compose build

# Step 2: Start the containers
docker compose up
```

### Api endpoints

1. **Create a User** : 
To create a user, make a POST request to user-service:

```bash
curl --location --request POST 'http://localhost:8081/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email" : "sachinthakan8coding@gmail.com"
}'
```


2. **Create a New Task** : To create a new task, use this POST request to task-manager:

```bash
curl --location --request POST 'http://localhost:8080/api/v1/tasks' \
--header 'Content-Type: application/json' \
--data-raw '{
  "title": "Finish project report2",
  "description": "Complete and submit the final project report by end of the week.",
  "priority": "high",
  "due_date": "2025-04-20",
  "user_id" : "U9d99f3aa51c446d98f3f233793ed552f"
}'
```


3. **Update a Task**: To update a task, use this PUT request to task-manager:

```bash
curl --location --request PUT 'http://localhost:8080/api/v1/tasks/T13c725341d0a4e84b86efc2991f103fe' \
--header 'Content-Type: application/json' \
--data-raw '{
  "title": "Finish project report",
  "description": "Complete and submit the final project report by end of the week.",
  "status" : "completed",
  "due_date": "2025-04-20",
  "user_id" : "U9d99f3aa51c446d98f3f233793ed552f"
}'
```


4. **Get Tasks (Paginated)**: To get paginated tasks, use this GET request:

```bash
curl --location --request GET 'http://localhost:8080/api/v1/tasks?user_id=U9d99f3aa51c446d98f3f233793ed552f&status=completed&limit=1'
```
