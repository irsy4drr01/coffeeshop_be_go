package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/handlers"
	middleware "github.com/irsy4drr01/coffeeshop_be_go/internal/middlewares"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
	"github.com/jmoiron/sqlx"
)

func product(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/product")

	repo := repositories.NewProduct(db)
	cld := pkg.NewCloudinaryUtil()

	handler := handlers.NewProduct(repo, cld)

	route.GET("/", handler.FetchAllProductsHandler)
	route.GET("/:uuid", handler.FetchDetailProductHandler)
	route.POST("/", middleware.AuthJwtMiddleware(), middleware.RoleAuthMiddleware("admin"), handler.PostProductHandler)
	route.PATCH("/:uuid", middleware.AuthJwtMiddleware(), middleware.RoleAuthMiddleware("admin"), handler.PatchProductHandler)
	route.DELETE("/:uuid", middleware.AuthJwtMiddleware(), middleware.RoleAuthMiddleware("admin"), handler.DeleteProductHandler)
}
