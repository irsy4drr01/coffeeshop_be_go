package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/handlers"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/services"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
	"github.com/jmoiron/sqlx"
)

func product(g *gin.RouterGroup, db *sqlx.DB) {
	route := g.Group("/products")

	repo := repositories.NewProduct(db)
	service := services.NewProductService(repo)
	cld := pkg.NewCloudinaryUtil()

	handler := handlers.NewProduct(service, cld)

	route.GET("/", handler.FetchAllProductsHandler)
	route.GET("/:id", handler.FetchProductDetailsHandler)
	// route.POST("/", middleware.AuthAndRoleMiddleware("admin"), handler.PostProductHandler)
	// route.PATCH("/:uuid", middleware.AuthAndRoleMiddleware("admin"), handler.PatchProductHandler)
	// route.DELETE("/:uuid", middleware.AuthAndRoleMiddleware("admin"), handler.DeleteProductHandler)
}
