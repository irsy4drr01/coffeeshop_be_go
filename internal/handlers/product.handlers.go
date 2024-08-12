package handlers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
)

type ProductHandlers struct {
	repo *repositories.RepoProduct
}

func NewProduct(r *repositories.RepoProduct) *ProductHandlers {
	return &ProductHandlers{repo: r}
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "uuid parameter is required"})
		return
	}

	var body map[string]any
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate only required fields for update
	product := models.Product{
		ProductName: body["product_name"].(string),
		Price:       body["price"].(int),
		Category:    body["category"].(string),
		Description: body["description"].(*string),
	}
	_, err := govalidator.ValidateStruct(product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
