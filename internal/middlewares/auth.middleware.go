package middleware

import (
	"log"
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

		log.Printf("Role from JWT: %s", check.Role)

		ctx.Set("userUuid", check.Uuid)
		ctx.Set("email", check.Email)
		ctx.Set("role", check.Role)
		ctx.Next()
	}
}

func RoleAuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Dapatkan role dari context
		role, exists := ctx.Get("role")
		if !exists {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden", "message": "Role not found"})
			ctx.Abort()
			return
		}

		log.Printf("Role in context: %s", role)

		// Cek role
		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				ctx.Next()
				return
			}
		}

		// Jika role tidak diizinkan
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden", "message": "You don't have the necessary permissions"})
		ctx.Abort()
	}
}
