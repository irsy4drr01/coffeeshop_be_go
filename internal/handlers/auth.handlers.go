package handlers

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
)

type AuthHandlers struct {
	repo repositories.AuthRepoInterface
}

func NewAuth(repo repositories.AuthRepoInterface) *AuthHandlers {
	return &AuthHandlers{repo: repo}
}

func (h *AuthHandlers) Register(ctx *gin.Context) {
	user := models.User{}

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Password, err = pkg.HashPassword(user.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password: " + err.Error()})
		return
	}

	response, createUser, err := h.repo.CreateUser(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"response": response, "data": createUser})
}

func (h *AuthHandlers) Login(ctx *gin.Context) {
	body := models.User{}

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "login failed", "message": err.Error()})
		return
	}

	_, err := govalidator.ValidateStruct(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "login failed", "message": err.Error()})
		return
	}

	result, err := h.repo.GetByEmail(body.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "login failed", "message": err.Error()})
		return
	}

	err = pkg.VerifyPassword(result.Password, body.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password", "message": err.Error()})
		return
	}

	jwt := pkg.NewJWT(result.Uuid, result.Email)
	token, err := jwt.GenerateToken()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "failed generate token", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "login success", "data": result, "token": token})
}