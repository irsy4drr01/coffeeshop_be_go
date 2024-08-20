package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/handlers"
	middleware "github.com/irsy4drr01/coffeeshop_be_go/internal/middlewares"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
	"github.com/jmoiron/sqlx"
)

func user(g *gin.Engine, db *sqlx.DB) {
	route := g.Group("/user")

	repo := repositories.NewUser(db)
	cld := pkg.NewCloudinaryUtil()

	handler := handlers.NewUser(repo, cld)

	route.GET("/", handler.FetchAllUserHandler)
	route.GET("/:uuid", middleware.AuthJwtMiddleware(), handler.FetchDetailUserHandler)
	route.PATCH("/:uuid", middleware.AuthJwtMiddleware(), handler.PatchUserHandler)
	route.DELETE("/:uuid", middleware.AuthJwtMiddleware(), handler.DeleteUserHandler)
}
