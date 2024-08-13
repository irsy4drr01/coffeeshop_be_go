package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
)

func AuthJwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var header string

		if header = ctx.GetHeader("Authorization"); header == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": "Authorization header missing"})
			ctx.Abort()
			return
		}

		if !strings.Contains(header, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": "Invalid Bearer Token"})
			ctx.Abort()
			return
		}

		// Bearer Bearer token
		token := strings.Replace(header, "Bearer ", "", -1)

		check, err := pkg.VerifyToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": "Invalid Bearer Token"})
			ctx.Abort()
			return
		}

		ctx.Set("userUuid", check.Uuid)
		ctx.Set("email", check.Email)
		ctx.Next()
	}
}
