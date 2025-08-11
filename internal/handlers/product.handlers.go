package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/services"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
)

type ProductHandlers struct {
	service services.ProductServiceInterface
	cld     pkg.CloudinaryInterface
}

func NewProduct(service services.ProductServiceInterface, cld pkg.CloudinaryInterface) *ProductHandlers {
	return &ProductHandlers{
		service: service,
		cld:     cld,
	}
}

// func (h *ProductHandlers) PostProductHandler(ctx *gin.Context) {
// 	responder := pkg.NewResponse(ctx)

// 	product := models.Product{}

// 	if err := ctx.ShouldBind(&product); err != nil {
// 		responder.BadRequest("Invalid request", err.Error())
// 		return
// 	}

// 	_, err := govalidator.ValidateStruct(product)
// 	if err != nil {
// 		responder.BadRequest("Validation failed", err.Error())
// 		return
// 	}

// 	createProduct, err := h.repo.CreateProduct(&product)
// 	if err != nil {
// 		responder.InternalServerError("Failed to create product", err.Error())
// 		return
// 	}

// 	responder.Created("Product created successfully.", createProduct)
// }

func (h *ProductHandlers) FetchAllProductsHandler(ctx *gin.Context) {
	var queryParams models.ProductQueryParams
	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		responder := pkg.NewResponse(ctx)
		responder.BadRequest("Invalid query params", err.Error())
		return
	}

	// Validasi sort by filter
	validSortBy := map[string]bool{
		"newest":     true,
		"oldest":     true,
		"asc":        true,
		"desc":       true,
		"cheapest":   true,
		"priciest":   true,
		"most_liked": true,
	}
	if !validSortBy[queryParams.SortBy] {
		queryParams.SortBy = "default"
	}

	// Build base URL and scheme
	scheme := "http://"
	if ctx.Request.TLS != nil {
		scheme = "https://"
	}
	baseHost := ctx.Request.Host
	baseURL := fmt.Sprintf("%s%s%s", scheme, baseHost, ctx.Request.URL.Path)

	// Call service
	response, meta, err := h.service.FetchAllProductsService(ctx.Request.Context(), queryParams, baseHost, scheme, baseURL)
	responder := pkg.NewResponse(ctx)
	if err != nil {
		ctx.Error(err)
		responder.InternalServerError("Internal server error", "Something went wrong, please try again later.")
		return
	}

	if len(response) == 0 {
		responder.SuccessWithMeta("No products found", response, meta)
		return
	}

	responder.SuccessWithMeta("Products fetched successfully", response, meta)
}

func (h *ProductHandlers) FetchProductDetailsHandler(ctx *gin.Context) {
	responder := pkg.NewResponse(ctx)

	productID := ctx.Param("id")
	if productID == "" {
		responder.BadRequest("Missing product ID", nil)
		return
	}

	// Build baseHost dan scheme di handler
	scheme := "http://"
	if ctx.Request.TLS != nil {
		scheme = "https://"
	}
	baseHost := ctx.Request.Host

	// Panggil service
	response, err := h.service.FetchProductDetailsService(ctx.Request.Context(), productID, baseHost, scheme)
	if err != nil {
		if err.Error() == "not found" {
			responder.NotFound("Product not found", nil)
		} else {
			ctx.Error(err)
			responder.InternalServerError("Internal server error", "Something went wrong, please try again later.")
		}
		return
	}

	responder.Success("Product detail fetched successfully", response)
}

// func (h *ProductHandlers) PatchProductHandler(ctx *gin.Context) {
// 	responder := pkg.NewResponse(ctx)

// 	uuid := ctx.Param("uuid")
// 	if uuid == "" {
// 		responder.BadRequest("UUID parameter is required", nil)
// 		return
// 	}

// 	// Ambil file gambar jika ada
// 	file, _, err := ctx.Request.FormFile("image")
// 	if err != nil && err != http.ErrMissingFile {
// 		responder.BadRequest("Failed to upload file", err.Error())
// 		return
// 	}

// 	// Ambil dari form-data (bukan dalam format JSON)
// 	productName := ctx.Request.FormValue("product_name")
// 	price := ctx.Request.FormValue("price")
// 	description := ctx.Request.FormValue("description")

// 	// Ambil data produk dari database
// 	existingProduct, err := h.repo.GetOneProduct(uuid)
// 	if err != nil {
// 		responder.InternalServerError("Failed to retrieve product", err.Error())
// 		return
// 	}

// 	// Validasi file jika ada
// 	if file != nil {
// 		buf := make([]byte, 512)
// 		if _, err := file.Read(buf); err != nil {
// 			responder.BadRequest("Failed to read file", err.Error())
// 			return
// 		}
// 		mimeType := http.DetectContentType(buf)

// 		// mimeType := header.Header.Get("Content-Type")
// 		if mimeType != "image/jpg" && mimeType != "image/jpeg" && mimeType != "image/png" {
// 			responder.BadRequest("Upload failed - wrong file type", nil)
// 			return
// 		}

// 		// Hapus file gambar lama di Cloudinary jika ada
// 		if existingProduct.Image != "" {
// 			publicID := pkg.GetPublicIDFromURL(existingProduct.Image)
// 			_, err := h.cld.DeleteFile(ctx, publicID)
// 			if err != nil {
// 				responder.InternalServerError("Failed to delete old file", err.Error())
// 				return
// 			}

// 			_, err = h.cld.DeleteFile(ctx, publicID)
// 			if err != nil {
// 				responder.InternalServerError("Failed to delete old file", err.Error())
// 				return
// 			}
// 		}

// 		// Upload file baru ke Cloudinary
// 		randomNumber, err := pkg.RandomInt(1000)
// 		if err != nil {
// 			responder.InternalServerError("Failed to generate random number", err.Error())
// 			return
// 		}

// 		fileName := fmt.Sprintf("product-image-%d", randomNumber)
// 		uploadResult, err := h.cld.UploadFile(ctx, file, fileName)
// 		if err != nil {
// 			responder.BadRequest("Failed to upload file", err.Error())
// 			return
// 		}

// 		// Set URL gambar di body
// 		existingProduct.Image = uploadResult.SecureURL
// 	}

// 	// Assign product attributes if they exist in the body
// 	if productName != "" {
// 		existingProduct.ProductName = productName
// 	}
// 	if price != "" {
// 		existingProduct.Price, _ = strconv.Atoi(price)
// 	}
// 	if description != "" {
// 		existingProduct.Description = &description
// 	} else if existingProduct.Description == nil {
// 		defaultDescription := ""
// 		existingProduct.Description = &defaultDescription
// 	}

// 	// Validasi Product
// 	if _, err := govalidator.ValidateStruct(existingProduct); err != nil {
// 		responder.BadRequest("Validation failed", err.Error())
// 		return
// 	}

// 	// Update product di database
// 	updatedProduct, err := h.repo.UpdateProduct(uuid, map[string]interface{}{
// 		"product_name": existingProduct.ProductName,
// 		"price":        existingProduct.Price,
// 		"description":  existingProduct.Description,
// 		"image":        existingProduct.Image,
// 	})

// 	if err != nil {
// 		responder.InternalServerError("Failed to update product", err.Error())
// 		return
// 	}

// 	responder.Success("Product updated successfully", updatedProduct)
// }

// ------------------------------------------

// func (h *ProductHandlers) DeleteProductHandler(ctx *gin.Context) {
// 	responder := pkg.NewResponse(ctx)

// 	uuid := ctx.Param("uuid")
// 	if uuid == "" {
// 		responder.BadRequest("UUID parameter is required", nil)
// 		return
// 	}

// 	deletedProduct, err := h.repo.DeleteProduct(uuid)
// 	if err != nil {
// 		responder.InternalServerError("Failed to delete product", err.Error())
// 		return
// 	}

// 	responder.Success("Product deleted successfully", deletedProduct)
// }

// ------------------------------------------
