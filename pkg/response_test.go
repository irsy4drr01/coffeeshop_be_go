package pkg

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	return c, w
}

func TestSuccess(t *testing.T) {
	c, w := setupTestContext()

	// Create a Responder instance
	responder := NewResponse(c)

	// Call Success method
	responder.Success("Success message", gin.H{"key": "value"})

	// Verify the response
	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody := `{"status":200,"message":"Success message","data":{"key":"value"}}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestLoginSuccess(t *testing.T) {
	c, w := setupTestContext()

	// Create a Responder instance
	responder := NewResponse(c)

	// Call LoginSuccess method
	token := "test-token"
	responder.LoginSuccess("Login successful", gin.H{"user": "testuser"}, token)

	// Verify the response
	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody := `{"status":200,"message":"Login successful","data":{"user":"testuser"},"token":"test-token"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestCreated(t *testing.T) {
	c, w := setupTestContext()

	// Create a Responder instance
	responder := NewResponse(c)

	// Call Created method
	responder.Created("Resource created", gin.H{"id": 1})

	// Verify the response
	assert.Equal(t, http.StatusCreated, w.Code)
	expectedBody := `{"status":201,"message":"Resource created","data":{"id":1}}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestBadRequest(t *testing.T) {
	c, w := setupTestContext()

	// Create a Responder instance
	responder := NewResponse(c)

	// Call BadRequest method
	responder.BadRequest("Bad request", "Invalid input")

	// Verify the response
	assert.Equal(t, http.StatusBadRequest, w.Code)
	expectedBody := `{"status":400,"message":"Bad request","error":"Invalid input"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestUnauthorized(t *testing.T) {
	c, w := setupTestContext()

	// Create a Responder instance
	responder := NewResponse(c)

	// Call Unauthorized method
	responder.Unauthorized("Unauthorized", "Invalid credentials")

	// Verify the response
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	expectedBody := `{"status":401,"message":"Unauthorized","error":"Invalid credentials"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestForbidden(t *testing.T) {
	c, w := setupTestContext()

	// Create a Responder instance
	responder := NewResponse(c)

	// Call Forbidden method
	responder.Forbidden("Forbidden", "No access rights")

	// Verify the response
	assert.Equal(t, http.StatusForbidden, w.Code)
	expectedBody := `{"status":403,"message":"Forbidden","error":"No access rights"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestNotFound(t *testing.T) {
	c, w := setupTestContext()

	// Create a Responder instance
	responder := NewResponse(c)

	// Call NotFound method
	responder.NotFound("Resource not found", "No data found")

	// Verify the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	expectedBody := `{"status":500,"message":"Resource not found","error":"No data found"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestInternalServerError(t *testing.T) {
	c, w := setupTestContext()

	// Create a Responder instance
	responder := NewResponse(c)

	// Call InternalServerError method
	responder.InternalServerError("Internal server error", "Something went wrong")

	// Verify the response
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	expectedBody := `{"status":500,"message":"Internal server error","error":"Something went wrong"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}
