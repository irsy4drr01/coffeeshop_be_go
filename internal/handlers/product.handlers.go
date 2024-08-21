package handlers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
)

type ProductHandlers struct {
	repo repositories.ProductRepoInterface
	cld  pkg.CloudinaryInterface
}

func NewProduct(repo repositories.ProductRepoInterface, cld pkg.CloudinaryInterface) *ProductHandlers {
	return &ProductHandlers{repo: repo, cld: cld}
}

func (h *ProductHandlers) PostProductHandler(ctx *gin.Context) {
	product := models.Product{}

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := govalidator.ValidateStruct(product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, createProduct, err := h.repo.CreateProduct(&product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"response": response, "data": createProduct})
}

func (h *ProductHandlers) FetchAllProductsHandler(ctx *gin.Context) {
	searchProductName := ctx.DefaultQuery("searchProductName", "")
	minPriceStr := ctx.DefaultQuery("minPrice", "0")
	maxPriceStr := ctx.DefaultQuery("maxPrice", strconv.Itoa(math.MaxInt32))
	category := ctx.DefaultQuery("category", "")
	sort := ctx.DefaultQuery("sort", "newest")

	// Pagination parameters
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")

	var minPrice, maxPrice, page, limit int
	var err error

	// Convert minPrice and maxPrice to int
	if minPrice, err = strconv.Atoi(minPriceStr); err != nil || minPrice < 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid minPrice"})
		return
	}
	if maxPriceStr == strconv.Itoa(math.MaxInt32) {
		maxPrice = math.MaxInt32 // Set to maximum value
	} else {
		if maxPrice, err = strconv.Atoi(maxPriceStr); err != nil || maxPrice < 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid maxPrice"})
			return
		}
	}

	// Convert page and limit to int
	if page, err = strconv.Atoi(pageStr); err != nil || page < 1 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}
	if limit, err = strconv.Atoi(limitStr); err != nil || limit < 1 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	data, err := h.repo.GetAllProducts(searchProductName, minPrice, maxPrice, category, sort, page, limit)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (h *ProductHandlers) FetchDetailProductHandler(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "uuid parameter is required"})
		return
	}

	product, err := h.repo.GetOneProduct(uuid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": product})
}

func (h *ProductHandlers) PatchProductHandler(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "UUID parameter is required"})
		return
	}

	// Ambil file gambar jika ada
	file, header, err := ctx.Request.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file: " + err.Error()})
		return
	}

	// Ambil JSON data dari form-data
	jsonStr := ctx.Request.FormValue("data")
	body := map[string]interface{}{}
	if jsonStr != "" {
		if err := json.Unmarshal([]byte(jsonStr), &body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
			return
		}
	}

	// Ambil data produk dari database
	existingProduct, err := h.repo.GetOneProduct(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product: " + err.Error()})
		return
	}

	// Validasi file jika ada
	if file != nil {
		mimeType := header.Header.Get("Content-Type")
		if mimeType != "image/jpg" && mimeType != "image/jpeg" && mimeType != "image/png" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Upload failed: wrong file type"})
			return
		}

		// Hapus file gambar lama di Cloudinary jika ada
		if existingProduct.Image != "" {
			publicID := pkg.GetPublicIDFromURL(existingProduct.Image)
			_, err := h.cld.DeleteFile(ctx, publicID)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete old file: " + err.Error()})
				return
			}
		}

		// Upload file baru ke Cloudinary
		randomNumber, err := pkg.RandomInt(1000)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate random number: " + err.Error()})
			return
		}

		fileName := fmt.Sprintf("product-image-%d", randomNumber)
		uploadResult, err := h.cld.UploadFile(ctx, file, fileName)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file: " + err.Error()})
			return
		}

		// Set URL gambar di body
		body["image"] = uploadResult.SecureURL
	}

	// Jika gambar tidak diupdate, gunakan gambar lama
	if _, exists := body["image"]; !exists || body["image"] == "" {
		body["image"] = existingProduct.Image
	}

	// Assign product attributes if they exist in the body
	product := models.Product{}
	if name, exists := body["product_name"].(string); exists && name != "" {
		product.ProductName = name
	}
	if price, exists := body["price"].(int); exists && price != 0 {
		product.Price = price
	}
	if category, exists := body["category"].(string); exists && category != "" {
		product.Category = category
	}
	if description, exists := body["description"].(string); exists && description != "" {
		product.Description = &description
	}

	// Validasi Product
	if _, err := govalidator.ValidateStruct(product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update product di database
	message, updatedProduct, err := h.repo.UpdateProduct(uuid, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": message, "data": updatedProduct})
}

func (h *ProductHandlers) DeleteProductHandler(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "uuid parameter is required"})
		return
	}

	message, deletedProduct, err := h.repo.DeleteProduct(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": message, "data": deletedProduct})
}
