package handlers

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"mime/multipart"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"testing"

// 	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
// 	"github.com/gin-gonic/gin"
// 	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
// 	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
// 	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func setupTestUser(userRepositoryMock *repositories.UserRepoMock, cloudinaryMock *pkg.CloudinaryMock) (*gin.Engine, *UserHandlers) {
// 	gin.SetMode(gin.TestMode)
// 	handlers := NewUser(userRepositoryMock, cloudinaryMock)
// 	router := gin.Default()
// 	return router, handlers
// }

// func TestFetchAllUsersHandler(t *testing.T) {
// 	userRepositoryMock := new(repositories.UserRepoMock)
// 	cloudinaryMock := new(pkg.CloudinaryMock)

// 	router, handlers := setupTestUser(userRepositoryMock, cloudinaryMock)

// 	router.GET("/user", handlers.FetchAllUserHandler)

// 	updatedAt1 := "2024-09-26T10:50:24.3399Z"
// 	updatedAt2 := "2024-09-26T10:53:30.3399Z"

// 	users := &models.Users{
// 		{
// 			Uuid:      "123e4567-e89b-12d3-a456-426614174000",
// 			Username:  "test user 1",
// 			Email:     "testuser1@example.com",
// 			Password:  "$2a$10$LMEF/rJp1mtbynxAap222u7zdBobjyFr8Sbepg2SnhqCT53mk022u",
// 			Image:     "http://example.com/userimage1.png",
// 			CreatedAt: "2024-09-11T20:18:28.159689Z",
// 			UpdatedAt: &updatedAt1,
// 		},
// 		{
// 			Uuid:      "a6dc7f8c-502e-4878-b78b-1cbf6195f2a7",
// 			Username:  "test user 2",
// 			Email:     "testuser2@example.com",
// 			Password:  "$2a$10$LMEF/rJp1mtbynxAap222u7zdBobjyFr8Sbepg2SnhqCT53mk022u",
// 			Image:     "http://example.com/userimage2.png",
// 			CreatedAt: "2024-09-11T20:20:00.159689Z",
// 			UpdatedAt: &updatedAt2,
// 		},
// 	}

// 	userRepositoryMock.On("GetAllUser", "", "", 1, 10).Return(users, nil)

// 	req, err := http.NewRequest(http.MethodGet, "/user", nil)
// 	assert.NoError(t, err, "An error occurred while making the request")

// 	recorder := httptest.NewRecorder()
// 	router.ServeHTTP(recorder, req)

// 	assert.Equal(t, http.StatusOK, recorder.Code, "Status code doesn't match")

// 	var actualResponse pkg.Response
// 	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
// 	assert.NoError(t, err, "Error: Failed get response")

// 	assert.Equal(t, 200, actualResponse.Status, "Status code doesn't match")
// 	assert.Equal(t, "Users fetched successfully", actualResponse.Message, "Message doesn't match")

// 	actualData := actualResponse.Data.([]interface{})

// 	assert.Len(t, actualData, 2, "Expected 2 items in the user list")

// 	user1 := actualData[0].(map[string]interface{})
// 	user2 := actualData[1].(map[string]interface{})

// 	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", user1["uuid"], "uuid doesn't match")
// 	assert.Equal(t, "test user 1", user1["username"], "username doesn't match")
// 	assert.Equal(t, "testuser1@example.com", user1["email"], "email doesn't match")
// 	assert.Equal(t, "$2a$10$LMEF/rJp1mtbynxAap222u7zdBobjyFr8Sbepg2SnhqCT53mk022u", user1["password"], "password doesn't match")
// 	assert.Equal(t, "http://example.com/userimage1.png", user1["image"], "image doesn't match")
// 	assert.Equal(t, "2024-09-11T20:18:28.159689Z", user1["created_at"], "created at doesn't match")
// 	assert.Equal(t, "2024-09-26T10:50:24.3399Z", user1["updated_at"], "updated at doesn't match")

// 	assert.Equal(t, "a6dc7f8c-502e-4878-b78b-1cbf6195f2a7", user2["uuid"], "uuid doesn't match")
// 	assert.Equal(t, "test user 2", user2["username"], "username doesn't match")
// 	assert.Equal(t, "testuser2@example.com", user2["email"], "email doesn't match")
// 	assert.Equal(t, "$2a$10$LMEF/rJp1mtbynxAap222u7zdBobjyFr8Sbepg2SnhqCT53mk022u", user2["password"], "password doesn't match")
// 	assert.Equal(t, "http://example.com/userimage2.png", user2["image"], "image doesn't match")
// 	assert.Equal(t, "2024-09-11T20:20:00.159689Z", user2["created_at"], "created at doesn't match")
// 	assert.Equal(t, "2024-09-26T10:53:30.3399Z", user2["updated_at"], "updated at doesn't match")

// 	userRepositoryMock.AssertExpectations(t)
// }

// func TestFetchDetailUserHandler(t *testing.T) {
// 	userRepositoryMock := new(repositories.UserRepoMock)
// 	cloudinaryMock := new(pkg.CloudinaryMock)

// 	router, handlers := setupTestUser(userRepositoryMock, cloudinaryMock)

// 	router.GET("/user/:uuid", handlers.FetchDetailUserHandler)

// 	user := models.User{
// 		Uuid:      "123e4567-e89b-12d3-a456-426614174000",
// 		Username:  "test user",
// 		Email:     "testuser@gmail.com",
// 		Password:  "$2a$10$LMEF/rJp1mtbynxAap222u7zdBobjyFr8Sbepg2SnhqCT53mk022u",
// 		Image:     "http://example.com/userimage.png",
// 		CreatedAt: "2024-09-11T20:18:28.159689Z",
// 	}

// 	userRepositoryMock.On("GetOneUser", "123e4567-e89b-12d3-a456-426614174000").Return(&user, nil)

// 	req, err := http.NewRequest("GET", "/user/123e4567-e89b-12d3-a456-426614174000", nil)
// 	assert.NoError(t, err, "An error occurred while making the request")

// 	recorder := httptest.NewRecorder()
// 	router.ServeHTTP(recorder, req)

// 	assert.Equal(t, http.StatusOK, recorder.Code, "Status code doesn't match")

// 	var actualResponse pkg.Response
// 	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
// 	assert.NoError(t, err, "Error: Failed get response")

// 	assert.Equal(t, 200, actualResponse.Status, "Status code doesn't match")
// 	assert.Equal(t, "User detail fetched successfully", actualResponse.Message, "Message doesn't match")

// 	actualData := actualResponse.Data.(map[string]interface{})

// 	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", actualData["uuid"], "uuid doesn't match")
// 	assert.Equal(t, "test user", actualData["username"], "user name doesn't match")
// 	assert.Equal(t, "testuser@gmail.com", actualData["email"], "email doesn't match")
// 	assert.Equal(t, "$2a$10$LMEF/rJp1mtbynxAap222u7zdBobjyFr8Sbepg2SnhqCT53mk022u", actualData["password"], "password doesn't match")
// 	assert.Equal(t, "http://example.com/userimage.png", actualData["image"], "image doesn't match")
// 	assert.Equal(t, "2024-09-11T20:18:28.159689Z", actualData["created_at"], "created_at doesn't match")

// 	userRepositoryMock.AssertExpectations(t)
// }

// func TestPatchUserHandler_Success(t *testing.T) {
// 	userRepositoryMock := new(repositories.UserRepoMock)
// 	cloudinaryMock := new(pkg.CloudinaryMock)

// 	router, handlers := setupTestUser(userRepositoryMock, cloudinaryMock)

// 	router.PATCH("/user/:uuid", handlers.PatchUserHandler)

// 	oldUpdatedAt := "2024-09-26T10:50:24.3399Z"
// 	user := models.User{
// 		Uuid:      "123e4567-e89b-12d3-a456-426614174000",
// 		Username:  "Old Username",
// 		Email:     "oldusername@example.com",
// 		Password:  "$2a$10$LMEF/rJp1mtbynxAap222u7zdBobjyFr8Sbepg2SnhqCT53mk022u",
// 		Image:     "http://example.com/old.png",
// 		UpdatedAt: &oldUpdatedAt,
// 	}

// 	userRepositoryMock.On("GetOneUser", "123e4567-e89b-12d3-a456-426614174000").Return(&user, nil)

// 	cloudinaryMock.On("UploadFile", mock.Anything, mock.Anything, mock.Anything).Return(&uploader.UploadResult{SecureURL: "http://example.com/new.png"}, nil)

// 	publicID := pkg.GetPublicIDFromURL(user.Image)
// 	cloudinaryMock.On("DeleteFile", mock.Anything, publicID).Return(&uploader.DestroyResult{}, nil)

// 	newUpdatedAt := "2024-09-26T10:55:38.3399Z"
// 	updatedUser := models.User{
// 		Uuid:      "123e4567-e89b-12d3-a456-426614174000",
// 		Username:  "New Username",
// 		Email:     "newusername@example.com",
// 		Password:  "$2a$10$LMEF/rJp1mtbynxAap222u7zdBobjyFr8Sbepg2SnhqCT53mk022u",
// 		Image:     "http://example.com/new.png",
// 		UpdatedAt: &newUpdatedAt,
// 	}

// 	userRepositoryMock.On("UpdateUser", "123e4567-e89b-12d3-a456-426614174000", mock.Anything).Return(&updatedUser, nil)

// 	form := &bytes.Buffer{}
// 	writer := multipart.NewWriter(form)
// 	_ = writer.WriteField("data", `{"product_name": "Updated Product", "price": 30000, "category": "Updated Category", "description": "Updated Description"}`)

// 	imagePath := "e:/BISMILLAH/2. Fazztrack/5. BootCamp/Week_12 (Beginer Backend Go)/Tugas/internal/testdata/img.png"
// 	fmt.Println("Image path:", imagePath) // Tambahkan ini untuk debugging

// 	file, err := os.Open(imagePath)
// 	if err != nil {
// 		t.Fatalf("Failed to open image file: %v", err)
// 	}
// 	defer file.Close()

// 	part, err := writer.CreateFormFile("image", "img.png")
// 	if err != nil {
// 		t.Fatalf("Failed to create form file: %v", err)
// 	}
// 	_, err = io.Copy(part, file)
// 	if err != nil {
// 		t.Fatalf("Failed to copy image file to form: %v", err)
// 	}

// 	writer.Close()

// 	req, err := http.NewRequest("PATCH", "/user/123e4567-e89b-12d3-a456-426614174000", form)
// 	assert.NoError(t, err, "An error occurred while making the request")
// 	req.Header.Set("Content-Type", writer.FormDataContentType())

// 	recorder := httptest.NewRecorder()
// 	router.ServeHTTP(recorder, req)

// 	assert.Equal(t, http.StatusOK, recorder.Code, "Status code doesn't match")

// 	var actualResponse pkg.Response
// 	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
// 	assert.NoError(t, err, "Error: Failed get response")

// 	assert.Equal(t, 200, actualResponse.Status, "Status code doesn't match")
// 	assert.Equal(t, "User updated successfully", actualResponse.Message, "Message doesn't match")

// 	actualData := actualResponse.Data.(map[string]interface{})

// 	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", actualData["uuid"], "uuid doesn't match")
// 	assert.Equal(t, "New Username", actualData["username"], "user name doesn't match")
// 	assert.Equal(t, "newusername@example.com", actualData["email"], "email doesn't match")
// 	assert.Equal(t, "$2a$10$LMEF/rJp1mtbynxAap222u7zdBobjyFr8Sbepg2SnhqCT53mk022u", actualData["password"], "password doesn't match")
// 	assert.Equal(t, "http://example.com/new.png", actualData["image"], "image doesn't match")
// 	assert.Equal(t, "2024-09-26T10:55:38.3399Z", actualData["updated_at"], "updated_at doesn't match")

// 	userRepositoryMock.AssertExpectations(t)
// 	cloudinaryMock.AssertExpectations(t)
// }

// func TestDeleteUserHandler(t *testing.T) {
// 	userRepositoryMock := new(repositories.UserRepoMock)
// 	cloudinaryMock := new(pkg.CloudinaryMock)

// 	router, handlers := setupTestUser(userRepositoryMock, cloudinaryMock)

// 	router.DELETE("/user/:uuid", handlers.DeleteUserHandler)

// 	user := &models.User{
// 		Uuid:      "123e4567-e89b-12d3-a456-426614174000",
// 		Username:  "Username",
// 		Email:     "username@example.com",
// 		IsDeleted: true,
// 	}

// 	userRepositoryMock.On("DeleteUser", "123e4567-e89b-12d3-a456-426614174000").Return(user, nil)

// 	req, err := http.NewRequest("DELETE", "/user/123e4567-e89b-12d3-a456-426614174000", nil)
// 	assert.NoError(t, err, "An error occurred while making the request")

// 	recorder := httptest.NewRecorder()
// 	router.ServeHTTP(recorder, req)

// 	assert.Equal(t, http.StatusOK, recorder.Code, "Status code doesn't match")

// 	var actualResponse pkg.Response
// 	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
// 	assert.NoError(t, err, "Error: Failed get response")

// 	assert.Equal(t, 200, actualResponse.Status, "Status code doesn't match")
// 	assert.Equal(t, "Delete user successfully", actualResponse.Message, "Message doesn't match")

// 	actualData := actualResponse.Data.(map[string]interface{})

// 	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", actualData["uuid"], "uuid doesn't match")
// 	assert.Equal(t, "Username", actualData["username"], "user name doesn't match")
// 	assert.Equal(t, "username@example.com", actualData["email"], "email doesn't match")
// 	assert.Equal(t, true, actualData["is_deleted"], "is_deleted doesn't match")

// 	userRepositoryMock.AssertExpectations(t)
// }
