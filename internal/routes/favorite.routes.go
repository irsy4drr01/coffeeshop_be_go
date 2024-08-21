package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/handlers"
	middleware "github.com/irsy4drr01/coffeeshop_be_go/internal/middlewares"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/jmoiron/sqlx"
)

func favorite(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/favorite")

	repo := repositories.NewFavorite(db)

	handler := handlers.NewFavorite(repo)

	route.GET("/:user_id", middleware.AuthJwtMiddleware(), middleware.RoleAuthMiddleware("user"), handler.GetFavoritesHandler)
	route.POST("/", middleware.AuthJwtMiddleware(), middleware.RoleAuthMiddleware("user"), handler.AddFavoriteHandler)
	route.DELETE("/:user_id/:product_id", middleware.AuthJwtMiddleware(), middleware.RoleAuthMiddleware("user"), handler.RemoveFavoriteHandler)
}
