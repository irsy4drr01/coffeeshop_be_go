package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/handlers"
	middleware "github.com/irsy4drr01/coffeeshop_be_go/internal/middlewares"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/services"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
	"github.com/jmoiron/sqlx"
)

func user(g *gin.RouterGroup, db *sqlx.DB) {
	repo := repositories.NewUser(db)
	service := services.NewUserService(repo)
	cld := pkg.NewCloudinaryUtil()
	handler := handlers.NewUser(service, cld)

	// routeUsers := g.Group("/users")
	// routeUsers.GET("/", middleware.AuthAndRoleMiddleware("admin"), handler.FetchAllUsersHandler)
	// routeUsers.GET("/:id", middleware.AuthAndRoleMiddleware("admin"), handler.FetchUserDetailsByAdminHandler)
	// routeUsers.PATCH("/:id", middleware.AuthAndRoleMiddleware("admin"), handler.PatchUserByAdminHandler)
	// routeUsers.DELETE("/:id", middleware.AuthAndRoleMiddleware("admin"), handler.DeleteUserByAdminHandler)

	routeUser := g.Group("/user")
	routeUser.GET("/profile", middleware.CORSMiddleware(), middleware.AuthAndRoleMiddleware(), handler.FetchProfileHandler)
	// routeUser.PATCH("/profile", middleware.CORSMiddleware(), middleware.AuthAndRoleMiddleware(), handler.PatchProfileService)
	// routeUser.PATCH("", middleware.AuthAndRoleMiddleware(), handler.PatchUserHandler)
	// routeUser.PATCH("/password", middleware.AuthAndRoleMiddleware(), handler.PatchPasswordHandler)
	// routeUser.DELETE("", middleware.AuthAndRoleMiddleware(), handler.DeleteUserHandler)

}
