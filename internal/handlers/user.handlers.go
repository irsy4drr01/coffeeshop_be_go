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

func (h *UserHandlers) PostUser(ctx *gin.Context) {
	user := models.User{}

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := h.CreateUser(&user)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, response)
}

func (h *UserHandlers) FetchAllUser(ctx *gin.Context) {
	data, err := h.GetAllUser()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(200, data)
}
