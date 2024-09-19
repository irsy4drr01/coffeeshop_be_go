package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
)

type FavoriteHandlers struct {
	repo repositories.FavoriteRepoInterface
}

func NewFavorite(repo repositories.FavoriteRepoInterface) *FavoriteHandlers {
	return &FavoriteHandlers{repo: repo}
}

func (h *FavoriteHandlers) AddFavoriteHandler(ctx *gin.Context) {
	responder := pkg.NewResponse(ctx)

	var req models.Favorite
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responder.BadRequest("Invalid request body", err.Error())
		return
	}

	favorite, err := h.repo.AddFavorite(req.UserID, req.ProductID)
	if err != nil {
		responder.InternalServerError("Failed to add product to favorites", err.Error())
		return
	}

	responder.Created("Product added to favorites successfully.", favorite)
}

func (h *FavoriteHandlers) RemoveFavoriteHandler(ctx *gin.Context) {
	responder := pkg.NewResponse(ctx)

	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		responder.BadRequest("Invalid user_id", err.Error())
		return
	}
	productID, err := strconv.Atoi(ctx.Param("product_id"))
	if err != nil {
		responder.BadRequest("Invalid product_id", err.Error())
		return
	}

	err = h.repo.RemoveFavorite(userID, productID)
	if err != nil {
		responder.InternalServerError("Failed to remove product from favorites", err.Error())
		return
	}

	responder.Success("Product removed from favorites successfully.", nil)
}

func (h *FavoriteHandlers) GetFavoritesHandler(ctx *gin.Context) {
	responder := pkg.NewResponse(ctx)

	userID, err := strconv.Atoi(ctx.Param("user_id"))
	if err != nil {
		responder.BadRequest("Invalid user_id", err.Error())
		return
	}

	favorites, err := h.repo.GetFavorites(userID)
	if err != nil {
		responder.InternalServerError("Failed to retrieve favorites", err.Error())
		return
	}

	responder.Success("Get favorites successfully.", favorites)
}
