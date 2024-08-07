package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/handlers"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/jmoiron/sqlx"
)

func product(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/product")

	repo := repositories.NewProduct(db)

	handler := handlers.NewProduct(repo)

	route.GET("/", handler.FetchAllProductsHandler)
	route.GET("/:uuid", handler.FetchDetailProductHandler)
	route.POST("/", handler.PostProductHandler)
	route.PATCH("/:uuid", handler.PatchProductHandler)
	route.DELETE("/:uuid", handler.DeleteProductHandler)
}
