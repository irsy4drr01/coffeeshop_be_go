package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/handlers"
	middleware "github.com/irsy4drr01/coffeeshop_be_go/internal/middlewares"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/services"
	"github.com/jmoiron/sqlx"
)

func orderHistory(g *gin.RouterGroup, db *sqlx.DB) {
	route := g.Group("/order-history")

	repo := repositories.NewOrderHistory(db)
	service := services.NewOrderHistoryService(repo)

	handler := handlers.NewOrderHistory(service)

	route.GET("/", middleware.AuthAndRoleMiddleware("user"), handler.FetchAllOrderHistoriesHandler)
	route.GET("/:order_id", middleware.AuthAndRoleMiddleware("user"), handler.FetchOrderHistoryDetailsHandler)
}
