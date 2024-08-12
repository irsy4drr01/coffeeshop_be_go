package handlers

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
)

type UserHandlers struct {
	repo repositories.UserRepoInterface
}

func NewUser(repo repositories.UserRepoInterface) *UserHandlers {
	return &UserHandlers{repo: repo}
}

func (h *UserHandlers) PostUserHandler(ctx *gin.Context) {
	user := models.User{}

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, createUser, err := h.repo.CreateUser(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": response, "data": createUser})
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

	var body map[string]any
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	user := models.User{}
	if username, exists := body["username"]; exists {
		user.Username = username.(string)
	}
	if email, exists := body["email"]; exists {
		user.Email = email.(string)
	}
	if password, exists := body["password"]; exists {
		user.Password = password.(string)
	}

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
