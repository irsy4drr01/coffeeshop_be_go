// auth/handlers/auth.handlers_test.go
package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	// Setup Gin router
	router := gin.Default()

	// Create mock repository
	authRepositoryMock := new(repositories.AuthRepoMock)

	// Create handler with mock repository
	handler := NewAuth(authRepositoryMock)

	// Define expected behavior
	authRepositoryMock.On("CreateUser", mock.Anything).Return(&models.User{
		Uuid:      "12345",
		Username:  "testuser",
		Email:     "testing@gmail.com",
		CreatedAt: "2024-01-01T00:00:00Z",
	}, nil)

	// Define the route
	router.POST("/auth/register", handler.Register)

	// Define request body
	requestBody, _ := json.Marshal(map[string]string{
		"username": "testuser",
		"email":    "testing@gmail.com",
		"password": "12345678",
	})

	// Create HTTP request
	req, err := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(requestBody))
	assert.NoError(t, err, "An error occurred while making the request")
	req.Header.Set("Content-Type", "application/json")

	// Create HTTP recorder
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, recorder.Code, "Status code doesn't match")

	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "Error: Failed get response")

	assert.Equal(t, 201, actualResponse.Status, "Status code doesn't match")
	assert.Equal(t, "User created successfully.", actualResponse.Message, "Message doesn't match")

	// var actualData map[string]interface{}
	// if data, ok := actualResponse.Data.(map[string]interface{}); ok {
	// 	actualData = data
	// } else {
	// 	t.Fatal("data is not in the expected format")
	// }

	actualData := actualResponse.Data.(map[string]interface{})

	assert.Equal(t, "12345", actualData["uuid"], "UUID doesn't match")
	assert.Equal(t, "testuser", actualData["username"], "Username doesn't match")
	assert.Equal(t, "testing@gmail.com", actualData["email"], "Email doesn't match")
	assert.Equal(t, "2024-01-01T00:00:00Z", actualData["created_at"], "CreatedAt doesn't match")
}

func TestLogin(t *testing.T) {
	router := gin.Default()
	authRepositoryMock := new(repositories.AuthRepoMock)
	handler := NewAuth(authRepositoryMock)

	authRepositoryMock.On("GetByEmail", mock.Anything).Return(&models.User{
		Uuid:     "12345",
		Username: "testuser",
		Email:    "testing@gmail.com",
		Password: "$2a$10$1rgEe.p59ssRcHmhqv..tel2nYyAx3XXLxPwRADNvydsMQueu6ybW",
	}, nil)

	router.POST("/auth/login", handler.Login)

	requestBody, _ := json.Marshal(map[string]string{
		"email":    "testing@gmail.com",
		"password": "password",
	})

	req, err := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(requestBody))
	assert.NoError(t, err, "An error occurred while making the request")
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Status code doesn't match")

	var actualResponse pkg.LoginResponse
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "Error: Failed get response")

	assert.Equal(t, 200, actualResponse.Status, "Status code doesn't match")
	assert.Equal(t, "Login success", actualResponse.Message, "Message doesn't match")

	// var actualData map[string]interface{}
	// if data, ok := actualResponse.Data.(map[string]interface{}); ok {
	// 	actualData = data
	// } else {
	// 	t.Fatal("data is not in the expected format")
	// }

	actualData := actualResponse.Data.(map[string]interface{})

	assert.Equal(t, "12345", actualData["uuid"], "UUID doesn't match")
	assert.Equal(t, "testuser", actualData["username"], "Username doesn't match")
	assert.Equal(t, "testing@gmail.com", actualData["email"], "Email doesn't match")

	assert.NotEmpty(t, actualResponse.Token, "Token should not be empty")
}
