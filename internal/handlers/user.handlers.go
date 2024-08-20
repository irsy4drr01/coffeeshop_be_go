package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
)

type UserHandlers struct {
	repo repositories.UserRepoInterface
	cld  pkg.CloudinaryInterface
}

func NewUser(repo repositories.UserRepoInterface, cld pkg.CloudinaryInterface) *UserHandlers {
	return &UserHandlers{repo: repo, cld: cld}
}

func (h *UserHandlers) FetchAllUserHandler(ctx *gin.Context) {
	searchUserName := ctx.DefaultQuery("searchUserName", "")
	sortBy := ctx.DefaultQuery("sort", "")

	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	users, err := h.repo.GetAllUser(searchUserName, sortBy, page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (h *UserHandlers) FetchDetailUserHandler(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "UUID parameter is required"})
		return
	}

	userDetail, err := h.repo.GetOneUser(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userDetail == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": userDetail})
}

func (h *UserHandlers) PatchUserHandler(ctx *gin.Context) {
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

	// Inisialisasi map body
	body := map[string]interface{}{}

	// Bind input dari form-data
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Ambil data pengguna dari database
	existingUser, err := h.repo.GetOneUser(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user: " + err.Error()})
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
		if existingUser.Image != "" {
			publicID := pkg.GetPublicIDFromURL(existingUser.Image)
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

		fileName := fmt.Sprintf("user-image-%d", randomNumber)
		uploadResult, err := h.cld.UploadFile(ctx, file, fileName)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file: " + err.Error()})
			return
		}

		// Set URL gambar di body
		body["image"] = uploadResult.SecureURL
	}

	// Assign user attributes if they exist in the body
	user := models.User{}
	if username, exists := body["username"].(string); exists && username != "" {
		user.Username = username
	}
	if email, exists := body["email"].(string); exists && email != "" {
		user.Email = email
	}
	if password, exists := body["password"].(string); exists && password != "" {
		user.Password = password
	}

	// Validasi User
	if _, err := govalidator.ValidateStruct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update user di database
	message, updatedUser, err := h.repo.UpdateUser(uuid, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": message, "data": updatedUser})
}

func (h *UserHandlers) DeleteUserHandler(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "UUID parameter is required"})
		return
	}

	message, deletedUser, err := h.repo.DeleteUser(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": message, "data": deletedUser})
}
