package accessor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"os"

)

var userServiceURL = os.Getenv("USER_SERVICE_API_URL")

// User represents a user object in your system
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// UserServiceAccessor is the struct for calling the user service APIs
type UserServiceAccessor struct {
	client *http.Client
}

// NewUserServiceAccessor creates a new instance of UserServiceAccessor
func NewUserServiceAccessor() *UserServiceAccessor {
	return &UserServiceAccessor{
		client: &http.Client{Timeout: 10 * time.Second}, // todo: make this configurable
	}
}

// GetUserByID fetches a user by their ID from the user service
func (a *UserServiceAccessor) GetUserByID(userID string) (*User, error) {
	url := fmt.Sprintf("%s/users/%s", userServiceURL, userID)
	resp, err := a.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	// Read and handle the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error response from server: %s", string(body))
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user data: %v", err)
	}

	return &user, nil
}

// CreateUser creates a new user via a POST request to the user service
func (a *UserServiceAccessor) CreateUser(newUser *User) (*User, error) {
	url := fmt.Sprintf("%s/users", userServiceURL)
	userData, err := json.Marshal(newUser)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user data: %v", err)
	}

	resp, err := a.client.Post(url, "application/json", bytes.NewBuffer(userData))
	if err != nil {
		return nil, fmt.Errorf("failed to make POST request: %v", err)
	}
	defer resp.Body.Close()

	// Read and handle the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("error response from server: %s", string(body))
	}

	var createdUser User
	if err := json.Unmarshal(body, &createdUser); err != nil {
		return nil, fmt.Errorf("failed to unmarshal created user data: %v", err) // todo: replace with logging
	}

	return &createdUser, nil
}