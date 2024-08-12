package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
)

type FavoriteHandlers struct {
	repo repositories.FavoriteRepoInterface
}

func NewFavorite(repo repositories.FavoriteRepoInterface) *FavoriteHandlers {
	return &FavoriteHandlers{repo: repo}
}

func (h *FavoriteHandlers) AddFavoriteHandler(ctx *gin.Context) {
	var req models.Favorite
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message, favorite, err := h.repo.AddFavorite(req.UserID, req.ProductID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": message, "data": favorite})
}

func (h *FavoriteHandlers) RemoveFavoriteHandler(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}
	productID, err := strconv.Atoi(ctx.Param("product_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product_id"})
		return
	}

	message, err := h.repo.RemoveFavorite(userID, productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": message})
}

func (h *FavoriteHandlers) GetFavoritesHandler(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	favorites, err := h.repo.GetFavorites(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": favorites})
}
