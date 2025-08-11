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
		// if !strings.Contains(header, "Bearer") {
		// 	responder.Unauthorized("Invalid Bearer Token", "Unauthorized")
		// 	ctx.Abort()
		// 	return
		// }

		// Ekstraksi token
		// token := strings.Replace(header, "Bearer ", "", -1)

		// Validasi format Bearer token [Versi Upgrade]
		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			responder.Unauthorized("Invalid Bearer Token", "Unauthorized")
			ctx.Abort()
			return
		}

		token := parts[1]

		// Verifikasi token
		check, err := pkg.VerifyToken(token)
		if err != nil {
			responder.Unauthorized("Invalid Bearer Token", "Unauthorized")
			ctx.Abort()
			return
		}

		// Simpan data dari token ke dalam context
		ctx.Set(pkg.ContextUserID, check.Uuid)
		ctx.Set(pkg.ContextEmail, check.Email)
		ctx.Set(pkg.ContextRole, check.Role)
		ctx.Set(pkg.ContextIsVerified, check.IsVerified)

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

// Verify Account Middleware (isVerified == true || isVerified == false)
func VerifyAccountMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		responder := pkg.NewResponse(ctx)

		isVerifiedAny, exists := ctx.Get(pkg.ContextIsVerified)
		if !exists {
			responder.Unauthorized("Verification status not found", "Unauthorized")
			ctx.Abort()
			return
		}

		isVerified, ok := isVerifiedAny.(bool)
		if !ok {
			responder.InternalServerError("Invalid verification status", "Server Error")
			ctx.Abort()
			return
		}

		if !isVerified {
			responder.Forbidden("Please verify your account to perform this action", "Forbidden")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
