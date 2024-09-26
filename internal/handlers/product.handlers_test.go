package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupTestProduct(productRepositoryMock *repositories.ProductRepoMock, cloudinaryMock *pkg.CloudinaryMock) (*gin.Engine, *ProductHandlers) {
	gin.SetMode(gin.TestMode)
	handlers := NewProduct(productRepositoryMock, cloudinaryMock)
	router := gin.Default()
	return router, handlers
}

func TestPostProductHandler(t *testing.T) {
	productRepositoryMock := new(repositories.ProductRepoMock)
	cloudinaryMock := new(pkg.CloudinaryMock)

	router, handlers := setupTestProduct(productRepositoryMock, cloudinaryMock)

	requestBody, _ := json.Marshal(map[string]interface{}{
		"product_name": "Test Product",
		"price":        20000,
		"category":     "Coffee",
		"description":  "Description 6",
	})

	router.POST("/products", handlers.PostProductHandler)

	description := "Test description"
	product := &models.Product{
		ID:          1,
		Uuid:        "12345",
		ProductName: "Test Product",
		Price:       10000,
		Category:    "Test Category",
		Image:       "",
		Description: &description,
		CreatedAt:   "2024-09-19T19:44:33.876019Z",
	}

	productRepositoryMock.On("CreateProduct", mock.Anything).Return(product, nil)

	req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(requestBody))
	assert.NoError(t, err, "An error occurred while making the request")
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code, "Status code doesn't match")

	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "Error: Failed get response")

	assert.Equal(t, 201, actualResponse.Status, "Status code doesn't match")
	assert.Equal(t, "Product created successfully.", actualResponse.Message, "Message doesn't match")

	actualData := actualResponse.Data.(map[string]interface{})

	assert.Equal(t, float64(1), actualData["id"], "user id doesn't match")
	assert.Equal(t, "12345", actualData["uuid"], "uuid doesn't match")
	assert.Equal(t, "Test Product", actualData["product_name"], "product name doesn't match")
	assert.Equal(t, float64(10000), actualData["price"], "price doesn't match")
	assert.Equal(t, "Test Category", actualData["category"], "category doesn't match")
	assert.Equal(t, "", actualData["image"], "image doesn't match")
	assert.Equal(t, "Test description", actualData["description"], "description doesn't match")
	assert.Equal(t, "2024-09-19T19:44:33.876019Z", actualData["created_at"], "created_at doesn't match")

	productRepositoryMock.AssertExpectations(t)
}

func TestFetchAllProductsHandler(t *testing.T) {
	productRepositoryMock := new(repositories.ProductRepoMock)
	cloudinaryMock := new(pkg.CloudinaryMock)

	router, handlers := setupTestProduct(productRepositoryMock, cloudinaryMock)

	router.GET("/products", handlers.FetchAllProductsHandler)

	description1 := "Description 1"
	description2 := "Description 2"
	updatedAt1 := "2024-08-21T05:02:02.847565Z"
	updatedAt2 := ""

	products := models.Products{
		{
			ID:          1,
			Uuid:        "12345",
			ProductName: "Product 1",
			Price:       10000,
			Category:    "Category1",
			Image:       "http://example.com/image1.png",
			Description: &description1,
			CreatedAt:   "2024-08-12T18:28:42.978583Z",
			UpdatedAt:   &updatedAt1,
		},
		{
			ID:          2,
			Uuid:        "12346",
			ProductName: "Product 2",
			Price:       15000,
			Category:    "Category1",
			Image:       "http://example.com/image2.png",
			Description: &description2,
			CreatedAt:   "2024-08-12T18:29:00.978583Z",
			UpdatedAt:   &updatedAt2,
		},
	}

	productRepositoryMock.On("GetAllProducts", "", 0, 2147483647, "", "newest", 1, 10).Return(&products, nil)

	req, err := http.NewRequest(http.MethodGet, "/products", nil)
	assert.NoError(t, err, "An error occurred while making the request")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Status code doesn't match")

	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "Error: Failed get response")

	assert.Equal(t, 200, actualResponse.Status, "Status code doesn't match")
	assert.Equal(t, "Products fetched successfully", actualResponse.Message, "Message doesn't match")

	actualData := actualResponse.Data.([]interface{})

	assert.Len(t, actualData, 2, "Expected 2 items in the product list")

	product1 := actualData[0].(map[string]interface{})
	product2 := actualData[1].(map[string]interface{})

	assert.Equal(t, float64(1), product1["id"], "id doesn't match")
	assert.Equal(t, "12345", product1["uuid"], "uuid doesn't match")
	assert.Equal(t, "Product 1", product1["product_name"], "product name doesn't match")
	assert.Equal(t, float64(10000), product1["price"], "price doesn't match")
	assert.Equal(t, "Category1", product1["category"], "category doesn't match")
	assert.Equal(t, "http://example.com/image1.png", product1["image"], "image doesn't match")
	assert.Equal(t, "Description 1", product1["description"], "description doesn't match")
	assert.Equal(t, "2024-08-12T18:28:42.978583Z", product1["created_at"], "created at doesn't match")
	assert.Equal(t, "2024-08-21T05:02:02.847565Z", product1["updated_at"], "updated at doesn't match")

	assert.Equal(t, float64(2), product2["id"], "id doesn't match")
	assert.Equal(t, "12346", product2["uuid"], "uuid doesn't match")
	assert.Equal(t, "Product 2", product2["product_name"], "product name doesn't match")
	assert.Equal(t, float64(15000), product2["price"], "price doesn't match")
	assert.Equal(t, "Category1", product2["category"], "category doesn't match")
	assert.Equal(t, "http://example.com/image2.png", product2["image"], "image doesn't match")
	assert.Equal(t, "Description 2", product2["description"], "description doesn't match")
	assert.Equal(t, "2024-08-12T18:29:00.978583Z", product2["created_at"], "created at doesn't match")
	assert.Equal(t, "", product2["updated_at"], "updated at doesn't match")

	productRepositoryMock.AssertExpectations(t)
}

func TestFetchDetailProductHandler(t *testing.T) {
	productRepositoryMock := new(repositories.ProductRepoMock)
	cloudinaryMock := new(pkg.CloudinaryMock)

	router, handlers := setupTestProduct(productRepositoryMock, cloudinaryMock)

	router.GET("/products/:uuid", handlers.FetchDetailProductHandler)

	description := "Test description"
	product := models.Product{
		ID:          1,
		Uuid:        "uuid1",
		ProductName: "Test Product",
		Price:       25000,
		Category:    "Test Category",
		Image:       "http://example.com/image1.png",
		Description: &description,
	}

	productRepositoryMock.On("GetOneProduct", "uuid10").Return(&product, nil)

	req, err := http.NewRequest(http.MethodGet, "/products/uuid10", nil)
	assert.NoError(t, err, "An error occurred while making the request")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Status code doesn't match")

	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "Error: Failed get response")

	assert.Equal(t, 200, actualResponse.Status, "Status code doesn't match")
	assert.Equal(t, "Product details fetched successfully", actualResponse.Message, "Message doesn't match")

	actualData := actualResponse.Data.(map[string]interface{})

	assert.Equal(t, float64(1), actualData["id"], "user id doesn't match")
	assert.Equal(t, "uuid1", actualData["uuid"], "uuid doesn't match")
	assert.Equal(t, "Test Product", actualData["product_name"], "product name doesn't match")
	assert.Equal(t, float64(25000), actualData["price"], "price doesn't match")
	assert.Equal(t, "Test Category", actualData["category"], "category doesn't match")
	assert.Equal(t, "http://example.com/image1.png", actualData["image"], "image doesn't match")
	assert.Equal(t, "Test description", actualData["description"], "description doesn't match")

	productRepositoryMock.AssertExpectations(t)
}

func TestPatchProductHandler(t *testing.T) {
	productRepositoryMock := new(repositories.ProductRepoMock)
	cloudinaryMock := new(pkg.CloudinaryMock)

	router, handlers := setupTestProduct(productRepositoryMock, cloudinaryMock)

	router.PATCH("/products/:uuid", handlers.PatchProductHandler)

	descriptionOld := "Old Description"
	descriptionUpdate := "Updated Description"

	existingProduct := models.Product{
		ID:          1,
		Uuid:        "41838fc7-3bc9-4895-8be0-5623fff20f59",
		ProductName: "Old Product",
		Price:       20000,
		Category:    "Old Category",
		Image:       "http://example.com/old.png",
		Description: &descriptionOld,
	}

	productRepositoryMock.On("GetOneProduct", "uuid1").Return(&existingProduct, nil)

	cloudinaryMock.On("UploadFile", mock.Anything, mock.Anything, mock.Anything).Return(&uploader.UploadResult{SecureURL: "http://example.com/new.png"}, nil)

	publicID := pkg.GetPublicIDFromURL(existingProduct.Image)
	cloudinaryMock.On("DeleteFile", mock.Anything, publicID).Return(&uploader.DestroyResult{}, nil)

	updatedAt := "2024-09-19T19:44:33.876019Z"
	productRepositoryMock.On("UpdateProduct", "uuid1", mock.Anything).Return(&models.Product{
		ID:          1,
		Uuid:        "41838fc7-3bc9-4895-8be0-5623fff20f59",
		ProductName: "Updated Product",
		Price:       30000,
		Category:    "Updated Category",
		Image:       "http://example.com/new.png",
		Description: &descriptionUpdate,
		UpdatedAt:   &updatedAt,
	}, nil)

	form := &bytes.Buffer{}
	writer := multipart.NewWriter(form)
	_ = writer.WriteField("data", `{"product_name": "Updated Product", "price": 30000, "category": "Updated Category", "description": "Updated Description"}`)

	imagePath := "e:/BISMILLAH/2. Fazztrack/5. BootCamp/Week_12 (Beginer Backend Go)/Tugas/internal/testdata/img.png"
	fmt.Println("Image path:", imagePath) // Tambahkan ini untuk debugging

	file, err := os.Open(imagePath)
	if err != nil {
		t.Fatalf("Failed to open image file: %v", err)
	}
	defer file.Close()

	part, err := writer.CreateFormFile("image", "img.png")
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatalf("Failed to copy image file to form: %v", err)
	}

	writer.Close()

	req, err := http.NewRequest(http.MethodPatch, "/products/uuid1", form)
	assert.NoError(t, err, "An error occurred while making the request")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Status code doesn't match")

	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "Error: Failed get response")

	assert.Equal(t, 200, actualResponse.Status, "Status code doesn't match")
	assert.Equal(t, "Product updated successfully", actualResponse.Message, "Message doesn't match")

	actualData := actualResponse.Data.(map[string]interface{})

	assert.Equal(t, float64(1), actualData["id"], "user id doesn't match")
	assert.Equal(t, "41838fc7-3bc9-4895-8be0-5623fff20f59", actualData["uuid"], "uuid doesn't match")
	assert.Equal(t, "Updated Product", actualData["product_name"], "product name doesn't match")
	assert.Equal(t, float64(30000), actualData["price"], "price doesn't match")
	assert.Equal(t, "Updated Category", actualData["category"], "category doesn't match")
	assert.Equal(t, "http://example.com/new.png", actualData["image"], "image doesn't match")
	assert.Equal(t, "Updated Description", actualData["description"], "description doesn't match")
	assert.Equal(t, "2024-09-19T19:44:33.876019Z", actualData["updated_at"], "updated_at doesn't match")

	productRepositoryMock.AssertExpectations(t)
	cloudinaryMock.AssertExpectations(t)
}

func TestDeleteProductHandler(t *testing.T) {
	productRepositoryMock := new(repositories.ProductRepoMock)
	cloudinaryMock := new(pkg.CloudinaryMock)

	router, handlers := setupTestProduct(productRepositoryMock, cloudinaryMock)

	router.DELETE("/product/:uuid", handlers.DeleteProductHandler)

	existingProduct := models.DeleteProduct{
		ID:          1,
		Uuid:        "uuid1",
		ProductName: "Test Product",
		IsDeleted:   true,
	}

	productRepositoryMock.On("DeleteProduct", "uuid1").Return(&existingProduct, nil)

	req, err := http.NewRequest(http.MethodDelete, "/product/uuid1", nil)
	assert.NoError(t, err, "An error occurred while making the request")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, "Status code doesn't match")

	var actualResponse pkg.Response
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponse)
	assert.NoError(t, err, "Error: Failed get response")

	assert.Equal(t, 200, actualResponse.Status, "Status code doesn't match")
	assert.Equal(t, "Product deleted successfully", actualResponse.Message, "Message doesn't match")

	actualData := actualResponse.Data.(map[string]interface{})

	assert.Equal(t, float64(1), actualData["id"], "user id doesn't match")
	assert.Equal(t, "uuid1", actualData["uuid"], "uuid doesn't match")
	assert.Equal(t, "Test Product", actualData["product_name"], "product name doesn't match")
	assert.Equal(t, true, actualData["is_deleted"], "is_deleted status doesn't match")

	productRepositoryMock.AssertExpectations(t)
}
