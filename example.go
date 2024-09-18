package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func example() {
	router := gin.Default()

	router.GET("/user", handler)
	router.GET("/user/:id", paramString) // reqParams or queryParams
	router.GET("/query", queryString)    // reqQuery
	router.POST("/user", reqBody)        // reqBody

	router.Run(":8000")
}

func handler(ctx *gin.Context) {
	ctx.String(200, "hello worlds")
}

func queryString(ctx *gin.Context) {
	page := ctx.Query("page")
	limit := ctx.Query("limit")

	ctx.JSON(200, gin.H{
		"page":  page,
		"limit": limit,
	})
}

func paramString(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(200, gin.H{
		"id": id,
	})
}

type User struct {
	Username string `form:"username" json:"username"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func reqBody(ctx *gin.Context) {
	var data User

	if err := ctx.ShouldBind(&data); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(200, data)
}
