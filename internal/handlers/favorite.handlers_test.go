package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
	"github.com/stretchr/testify/assert"
)

func setupTestFavorite() (*gin.Engine, *repositories.FavoriteRepoMock, *FavoriteHandlers) {
	gin.SetMode(gin.TestMode)
	favoriteRepositoryMock := new(repositories.FavoriteRepoMock)
	handlers := NewFavorite(favoriteRepositoryMock)

	r := gin.Default()
	return r, favoriteRepositoryMock, handlers
}

func TestAddFavoriteHandler(t *testing.T) {
	r, favoriteRepositoryMock, handlers := setupTestFavorite()

	// Prepare request body
	UserID := 1
	ProductID := 1
	requestBody, _ := json.Marshal(map[string]int{
		"user_id":    UserID,
		"product_id": ProductID,
	})

	// Define the route
	r.POST("/favorite", handlers.AddFavoriteHandler)

	// Mock the AddFavorite method
	favoriteRepositoryMock.On("AddFavorite", UserID, ProductID).Return(&models.Favorite{
		UserID:      UserID,
		Username:    "testuser",
		ProductID:   ProductID,
		ProductName: "Chicken Sandwich",
	}, nil)

	// Create request and response
	req, err := http.NewRequest(http.MethodPost, "/favorite", bytes.NewBuffer(requestBody))
	assert.NoError(t, err, "An error occurred while making the request")
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	// Check response
	assert.Equal(t, http.StatusCreated, recorder.Code, "Status code doesn't match")

	// Parse response body
	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "Error: Failed get response")

	// Verify response content
	assert.Equal(t, 201, actualResponse.Status, "Status code doesn't match")
	assert.Equal(t, "Product added to favorites successfully.", actualResponse.Message, "Message doesn't match")

	actualData := actualResponse.Data.(map[string]interface{})

	assert.Equal(t, float64(1), actualData["user_id"], "user id doesn't match")
	assert.Equal(t, "testuser", actualData["username"], "username doesn't match")
	assert.Equal(t, float64(1), actualData["product_id"], "product id doesn't match")
	assert.Equal(t, "Chicken Sandwich", actualData["product_name"], "product name doesn't match")

	// Ensure the mock expectations were met
	favoriteRepositoryMock.AssertExpectations(t)
}

func TestRemoveFavoriteHandler(t *testing.T) {
	r, favoriteRepositoryMock, handlers := setupTestFavorite()

	userID := 1
	productID := 1

	r.DELETE("/favorite/:user_id/:product_id", handlers.RemoveFavoriteHandler)

	favoriteRepositoryMock.On("RemoveFavorite", userID, productID).Return(nil)

	req, err := http.NewRequest("DELETE", "/favorite/"+strconv.Itoa(userID)+"/"+strconv.Itoa(productID), nil)
	assert.NoError(t, err, "An error occurred while making the request")

	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	// Check response
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Parse response body
	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "Error: Failed get response")

	// Verify response content
	assert.Equal(t, "Product removed from favorites successfully.", actualResponse.Message, "Message doesn't match")

	// Ensure the mock expectations were met
	favoriteRepositoryMock.AssertExpectations(t)
}

func TestGetFavoritesHandler(t *testing.T) {
	r, favoriteRepositoryMock, handlers := setupTestFavorite()

	r.GET("/favorite/:user_id", handlers.GetFavoritesHandler)

	userID := 1

	mockFavorites := &models.Favorites{
		{
			UserID:      userID,
			Username:    "testuser1",
			ProductID:   1,
			ProductName: "Chicken Sandwich",
		},
		{
			UserID:      userID,
			Username:    "testuser1",
			ProductID:   2,
			ProductName: "Kentang Goreng",
		},
	}

	favoriteRepositoryMock.On("GetFavorites", userID).Return(mockFavorites, nil)

	req, err := http.NewRequest(http.MethodGet, "/favorite/"+strconv.Itoa(userID), nil)
	assert.NoError(t, err, "An error occurred while making the request")

	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Status code doesn't match")

	// Parse response body
	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "Error: Failed get response")

	// Verify response content
	assert.Equal(t, "Get favorites successfully.", actualResponse.Message, "Message doesn't match")

	// Check the data returned
	actualData := actualResponse.Data.([]interface{})

	assert.Len(t, actualData, 2, "Expected 2 items in the favorites list")
	assert.Equal(t, "Chicken Sandwich", actualData[0].(map[string]interface{})["product_name"], "product name doesn't match")
	assert.Equal(t, "Kentang Goreng", actualData[1].(map[string]interface{})["product_name"], "product name doesn't match")

	// Ensure the mock expectations were met
	favoriteRepositoryMock.AssertExpectations(t)
}
