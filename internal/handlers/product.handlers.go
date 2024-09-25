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
	responder := pkg.NewResponse(ctx)

	product := models.Product{}

	if err := ctx.ShouldBind(&product); err != nil {
		responder.BadRequest("Invalid request", err.Error())
		return
	}

	_, err := govalidator.ValidateStruct(product)
	if err != nil {
		responder.BadRequest("Validation failed", err.Error())
		return
	}

	createProduct, err := h.repo.CreateProduct(&product)
	if err != nil {
		responder.InternalServerError("Failed to create product", err.Error())
		return
	}

	responder.Created("Product created successfully.", createProduct)
}

func (h *ProductHandlers) FetchAllProductsHandler(ctx *gin.Context) {
	responder := pkg.NewResponse(ctx)

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
		responder.BadRequest("Invalid minPrice", err.Error())
		return
	}
	if maxPriceStr == strconv.Itoa(math.MaxInt32) {
		maxPrice = math.MaxInt32 // Set to maximum value
	} else {
		if maxPrice, err = strconv.Atoi(maxPriceStr); err != nil || maxPrice < 0 {
			responder.BadRequest("Invalid maxPrice", err.Error())
			return
		}
	}

	// Convert page and limit to int
	if page, err = strconv.Atoi(pageStr); err != nil || page < 1 {
		responder.BadRequest("Invalid page number", err.Error())
		return
	}

	if limit, err = strconv.Atoi(limitStr); err != nil || limit < 1 {
		responder.BadRequest("Invalid limit", err.Error())
		return
	}

	data, err := h.repo.GetAllProducts(searchProductName, minPrice, maxPrice, category, sort, page, limit)
	if err != nil {
		responder.InternalServerError("Failed to fetch products", err.Error())
		return
	}

	responder.Success("Products fetched successfully", data)
}

func (h *ProductHandlers) FetchDetailProductHandler(ctx *gin.Context) {
	responder := pkg.NewResponse(ctx)

	uuid := ctx.Param("uuid")
	if uuid == "" {
		responder.BadRequest("UUID parameter is required", nil)
		return
	}

	product, err := h.repo.GetOneProduct(uuid)
	if err != nil {
		responder.NotFound("Product not found", err.Error())
		return
	}

	responder.Success("Product details fetched successfully", product)
}

func (h *ProductHandlers) PatchProductHandler(ctx *gin.Context) {
	responder := pkg.NewResponse(ctx)

	uuid := ctx.Param("uuid")
	if uuid == "" {
		responder.BadRequest("UUID parameter is required", nil)
		return
	}

	// Ambil file gambar jika ada
	file, _, err := ctx.Request.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		responder.BadRequest("Failed to upload file", err.Error())
		return
	}

	// Ambil JSON data dari form-data
	jsonStr := ctx.Request.FormValue("data")
	body := map[string]interface{}{}
	if jsonStr != "" {
		if err := json.Unmarshal([]byte(jsonStr), &body); err != nil {
			responder.BadRequest("Failed to upload file", err.Error())
			return
		}
	}

	// Ambil data produk dari database
	existingProduct, err := h.repo.GetOneProduct(uuid)
	if err != nil {
		responder.InternalServerError("Failed to retrieve product", err.Error())
		return
	}

	// Validasi file jika ada
	if file != nil {
		buf := make([]byte, 512)
		if _, err := file.Read(buf); err != nil {
			responder.BadRequest("Failed to read file", err.Error())
			return
		}
		mimeType := http.DetectContentType(buf)

		// mimeType := header.Header.Get("Content-Type")
		if mimeType != "image/jpg" && mimeType != "image/jpeg" && mimeType != "image/png" {
			responder.BadRequest("Upload failed - wrong file type", nil)
			return
		}

		// Hapus file gambar lama di Cloudinary jika ada
		if existingProduct.Image != "" {
			publicID := pkg.GetPublicIDFromURL(existingProduct.Image)
			_, err := h.cld.DeleteFile(ctx, publicID)
			if err != nil {
				responder.InternalServerError("Failed to delete old file", err.Error())
				return
			}

			_, err = h.cld.DeleteFile(ctx, publicID)
			if err != nil {
				responder.InternalServerError("Failed to delete old file", err.Error())
				return
			}
		}

		// Upload file baru ke Cloudinary
		randomNumber, err := pkg.RandomInt(1000)
		if err != nil {
			responder.InternalServerError("Failed to generate random number", err.Error())
			return
		}

		fileName := fmt.Sprintf("product-image-%d", randomNumber)
		uploadResult, err := h.cld.UploadFile(ctx, file, fileName)
		if err != nil {
			responder.BadRequest("Failed to upload file", err.Error())
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
		responder.BadRequest("Validation failed", err.Error())
		return
	}

	// Update product di database
	updatedProduct, err := h.repo.UpdateProduct(uuid, body)
	if err != nil {
		responder.InternalServerError("Failed to update product", err.Error())
		return
	}

	responder.Success("Product updated successfully", updatedProduct)
}

func (h *ProductHandlers) DeleteProductHandler(ctx *gin.Context) {
	responder := pkg.NewResponse(ctx)

	uuid := ctx.Param("uuid")
	if uuid == "" {
		responder.BadRequest("UUID parameter is required", nil)
		return
	}

	deletedProduct, err := h.repo.DeleteProduct(uuid)
	if err != nil {
		responder.InternalServerError("Failed to delete product", err.Error())
		return
	}

	responder.Success("Product deleted successfully", deletedProduct)
}
