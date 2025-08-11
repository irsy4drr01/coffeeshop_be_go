package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	mainRouter := router.Group("/api")

	auth(mainRouter, db)
	user(mainRouter, db)
	product(mainRouter, db)
	order(mainRouter, db)
	orderHistory(mainRouter, db)
	favorite(mainRouter, db)

	return router
}
