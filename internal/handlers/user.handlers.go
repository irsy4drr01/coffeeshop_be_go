package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
)

type UserHandlers struct {
	*repositories.RepoUser
}

func NewUser(r *repositories.RepoUser) *UserHandlers {
	return &UserHandlers{r}
}

func (h *UserHandlers) PostUserHandler(ctx *gin.Context) {
	user := models.User{}

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, createUser, err := h.CreateUser(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"response": response, "data": createUser})
}

func (h *UserHandlers) FetchAllUserHandler(ctx *gin.Context) {
	searchUserName := ctx.Query("searchUserName")
	sortBy := ctx.Query("sort")

	data, err := h.GetAllUser(searchUserName, sortBy)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (h *UserHandlers) PatchUserHandler(ctx *gin.Context) {
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

	message, updatedUser, err := h.UpdateUser(uuid, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": message, "data": updatedUser})
}

func (h *UserHandlers) DeleteUserHandler(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "uuid parameter is required"})
		return
	}

	message, deletedUser, err := h.DeleteUser(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": message, "data": deletedUser})
}
