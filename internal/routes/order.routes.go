package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/handlers"
	middleware "github.com/irsy4drr01/coffeeshop_be_go/internal/middlewares"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/services"
	"github.com/jmoiron/sqlx"
)

func order(g *gin.RouterGroup, db *sqlx.DB) {
	route := g.Group("/order")

	repo := repositories.NewOrder(db)
	service := services.NewOrderService(repo)

	handler := handlers.NewOrder(service)

	route.POST("/", middleware.AuthAndRoleMiddleware("user"), handler.AddOrderHandler)
}
