package middleware

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
)

func AuthAndRoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		responder := pkg.NewResponse(ctx)

		// Pengecekan header Authorization
		header := ctx.GetHeader("Authorization")
		if header == "" {
			responder.Unauthorized("Authorization header missing", "Unauthorized")
			ctx.Abort()
			return
		}

		// Validasi format Bearer token
		if !strings.Contains(header, "Bearer") {
			responder.Unauthorized("Invalid Bearer Token", "Unauthorized")
			ctx.Abort()
			return
		}

		// Ekstraksi token
		token := strings.Replace(header, "Bearer ", "", -1)

		// Verifikasi token
		check, err := pkg.VerifyToken(token)
		if err != nil {
			responder.Unauthorized("Invalid Bearer Token", "Unauthorized")
			ctx.Abort()
			return
		}

		// Simpan data dari token ke dalam context
		ctx.Set("userUuid", check.Uuid)
		ctx.Set("email", check.Email)
		ctx.Set("role", check.Role)

		// Pengecekan role
		if len(allowedRoles) > 0 {
			role := check.Role
			log.Printf("Role from JWT: %s", role)

			// Cek apakah role diperbolehkan
			roleAllowed := false
			for _, allowedRole := range allowedRoles {
				if role == allowedRole {
					roleAllowed = true
					break
				}
			}

			// Jika role tidak diizinkan
			if !roleAllowed {
				responder.Forbidden("You don't have the necessary permissions", "Forbidden")
				ctx.Abort()
				return
			}
		}

		ctx.Next()
	}
}
