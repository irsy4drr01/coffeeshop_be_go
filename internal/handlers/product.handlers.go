package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
)

type ProductHandlers struct {
	*repositories.RepoProduct
}

func NewProduct(r *repositories.RepoProduct) *ProductHandlers {
	return &ProductHandlers{r}
}

func (h *ProductHandlers) PostProductHandler(ctx *gin.Context) {
	product := models.Product{}

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, createProduct, err := h.CreateProduct(&product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"response": response, "data": createProduct})
}

func (h *ProductHandlers) FetchAllProductsHandler(ctx *gin.Context) {
	searchProductName := ctx.Query("searchProductName")
	minPriceStr := ctx.Query("minPrice")
	maxPriceStr := ctx.Query("maxPrice")
	category := ctx.Query("category")
	sort := ctx.Query("sort")

	// Pagination parameters
	pageStr := ctx.Query("page")
	limitStr := ctx.Query("limit")

	var minPrice, maxPrice, page, limit int
	var err error

	// Convert minPrice and maxPrice to int
	if minPriceStr != "" {
		if minPrice, err = strconv.Atoi(minPriceStr); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid minPrice"})
			return
		}
	}
	if maxPriceStr != "" {
		if maxPrice, err = strconv.Atoi(maxPriceStr); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid maxPrice"})
			return
		}
	}

	// Convert page and limit to int
	if pageStr == "" {
		page = 1
	} else {
		if page, err = strconv.Atoi(pageStr); err != nil || page < 1 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
			return
		}
	}
	if limitStr == "" {
		limit = 10 // Default limit
	} else {
		if limit, err = strconv.Atoi(limitStr); err != nil || limit < 1 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			return
		}
	}

	data, err := h.GetAllProducts(searchProductName, minPrice, maxPrice, category, sort, page, limit)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
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

	message, updatedProduct, err := h.UpdateProduct(uuid, body)
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

	message, deletedProduct, err := h.DeleteProduct(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": message, "data": deletedProduct})
}
