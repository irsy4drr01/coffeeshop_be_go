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
	minPrice := ctx.Query("minPrice")
	maxPrice := ctx.Query("maxPrice")
	category := ctx.Query("category")
	sort := ctx.Query("sort")

	// Convert minPrice and maxPrice to float64
	var minPriceFloat, maxPriceFloat float64
	if minPrice != "" {
		if mp, err := strconv.ParseFloat(minPrice, 64); err == nil {
			minPriceFloat = mp
		}
	}
	if maxPrice != "" {
		if mp, err := strconv.ParseFloat(maxPrice, 64); err == nil {
			maxPriceFloat = mp
		}
	}

	data, err := h.GetAllProducts(searchProductName, minPriceFloat, maxPriceFloat, category, sort)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
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
