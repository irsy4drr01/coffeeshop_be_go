package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func New(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	auth(router, db)
	user(router, db)
	product(router, db)
	favorite(router, db)

	return router
}
