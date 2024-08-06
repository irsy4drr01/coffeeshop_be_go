package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/handlers"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/jmoiron/sqlx"
)

func user(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/user")

	repo := repositories.NewUser(db)

	handler := handlers.NewUser(repo)

	route.GET("/", handler.FetchAllUserHandler)
	route.POST("/", handler.PostUserHandler)
	route.PATCH("/:uuid", handler.PatchUserHandler)
	route.DELETE("/:uuid", handler.DeleteUserHandler)
}
