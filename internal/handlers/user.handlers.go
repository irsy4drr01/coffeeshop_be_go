package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
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
	responder := pkg.NewResponse(ctx)

	searchUserName := ctx.DefaultQuery("searchUserName", "")
	sortBy := ctx.DefaultQuery("sort", "")

	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		responder.BadRequest("Error", "Invalid page parameter")
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		responder.BadRequest("Invalid limit parameter", err.Error())
		return
	}

	result, err := h.repo.GetAllUser(searchUserName, sortBy, page, limit)
	if err != nil {
		responder.InternalServerError("Failed to fetch users", err.Error())
		return
	}

	responder.Success("Users fetched successfully", result)
}

func (h *UserHandlers) FetchDetailUserHandler(ctx *gin.Context) {
	responder := pkg.NewResponse(ctx)

	uuid := ctx.Param("uuid")
	if uuid == "" {
		responder.BadRequest("UUID parameter is required", nil)
		return
	}

	userDetail, err := h.repo.GetOneUser(uuid)
	if err != nil {
		responder.InternalServerError("Failed to fetch user detail", err.Error())
		return
	}

	if userDetail == nil {
		responder.NotFound("User not found", nil)
		return
	}

	responder.Success("User detail fetched successfully", userDetail)
}

func (h *UserHandlers) PatchUserHandler(ctx *gin.Context) {
	responder := pkg.NewResponse(ctx)

	uuid := ctx.Param("uuid")
	if uuid == "" {
		responder.BadRequest("UUID parameter is required", nil)
		return
	}

	// Ambil file gambar jika ada
	file, header, err := ctx.Request.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		responder.BadRequest("Failed to upload file", err.Error())
		return
	}

	// Ambil data dari form-data (bukan dalam format JSON lagi)
	username := ctx.Request.FormValue("username")
	email := ctx.Request.FormValue("email")
	password := ctx.Request.FormValue("password")

	// Ambil data pengguna dari database
	user, err := h.repo.GetOneUser(uuid)
	if err != nil {
		responder.InternalServerError("Failed to retrieve user", err.Error())
		return
	}

	// Validasi file jika ada
	if file != nil {
		mimeType := header.Header.Get("Content-Type")
		if mimeType != "image/jpg" && mimeType != "image/jpeg" && mimeType != "image/png" {
			responder.BadRequest("Upload failed - wrong file type", nil)
			return
		}

		// Hapus file gambar lama di Cloudinary jika ada
		if user.Image != "" {
			publicID := pkg.GetPublicIDFromURL(user.Image)
			_, err := h.cld.DeleteFile(ctx, publicID)
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

		fileName := fmt.Sprintf("user-image-%d", randomNumber)
		uploadResult, err := h.cld.UploadFile(ctx, file, fileName)
		if err != nil {
			responder.BadRequest("Failed to upload file", err.Error())
			return
		}

		// Set URL gambar di body
		user.Image = uploadResult.SecureURL
	}

	// Hash password jika ada
	if password != "" {
		hashedPassword, err := pkg.HashPassword(password)
		if err != nil {
			responder.InternalServerError("Failed to hash password", err.Error())
			return
		}
		user.Password = hashedPassword
	}

	// Assign user attributes if they exist in the body
	if username != "" {
		user.Username = username
	}
	if email != "" {
		user.Email = email
	}

	// Validasi User
	if _, err := govalidator.ValidateStruct(user); err != nil {
		responder.BadRequest("Validation error", err.Error())
		return
	}

	// Update user di database
	updatedUser, err := h.repo.UpdateUser(uuid, map[string]interface{}{
		"username": user.Username,
		"email":    user.Email,
		"password": user.Password,
		"image":    user.Image,
	})

	if err != nil {
		responder.InternalServerError("Failed to update user", err.Error())
		return
	}

	responder.Success("User updated successfully", updatedUser)
}

func (h *UserHandlers) DeleteUserHandler(ctx *gin.Context) {
	responder := pkg.NewResponse(ctx)

	uuid := ctx.Param("uuid")
	if uuid == "" {
		responder.BadRequest("UUID parameter is required", nil)
		return
	}

	deletedUser, err := h.repo.DeleteUser(uuid)
	if err != nil {
		responder.InternalServerError("Failed to delete user", err.Error())
		return
	}

	if deletedUser == nil {
		responder.NotFound("User not found", nil)
		return
	}

	responder.Success("Delete user successfully", deletedUser)
}
