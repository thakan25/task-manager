# Use Golang base image
FROM golang:1.20-alpine as builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the application
RUN go build -o task-manager .

# Use a smaller base image for the final container
FROM alpine:latest

# Install necessary dependencies (if needed)
RUN apk add --no-cache ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/task-manager .

# Expose the port the app will run on
EXPOSE 8080

# Command to run the application
CMD ["./task-manager"]
