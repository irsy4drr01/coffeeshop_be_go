package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/handlers"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/jmoiron/sqlx"
)

func auth(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/auth")

	repo := repositories.NewAuth(db)

	handler := handlers.NewAuth(repo)

	route.POST("/register", handler.Register)
	route.POST("/login", handler.Login)
}
