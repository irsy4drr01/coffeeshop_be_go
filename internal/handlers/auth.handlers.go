package handlers

import (
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
	responder := pkg.NewResponse(ctx)

	user := models.User{}

	if err := ctx.ShouldBind(&user); err != nil {
		responder.BadRequest("Invalid input", err.Error())
		return
	}

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		responder.BadRequest("Validation failed", err.Error())
		return
	}

	user.Password, err = pkg.HashPassword(user.Password)
	if err != nil {
		responder.InternalServerError("Failed to hash password", err.Error())
		return
	}

	createUser, err := h.repo.CreateUser(&user)
	if err != nil {
		responder.InternalServerError("Internal Server Error", err.Error())
		return
	}

	data := models.User{
		Uuid:      createUser.Uuid,
		Username:  createUser.Username,
		Email:     createUser.Email,
		CreatedAt: createUser.CreatedAt,
	}

	responder.Created("User created successfully.", data)
}

func (h *AuthHandlers) Login(ctx *gin.Context) {
	responder := pkg.NewResponse(ctx)

	body := models.User{}

	if err := ctx.ShouldBind(&body); err != nil {
		responder.BadRequest("Login failed", err.Error())
		return
	}

	_, err := govalidator.ValidateStruct(&body)
	if err != nil {
		responder.BadRequest("Login failed!", err.Error())
		return
	}

	data, err := h.repo.GetByEmail(body.Email)
	if err != nil {
		responder.BadRequest("Login failed!", err.Error())
		return
	}

	err = pkg.VerifyPassword(data.Password, body.Password)
	if err != nil {
		responder.BadRequest("Wrong password!", err.Error())
		return
	}

	jwt := pkg.NewJWT(data.Uuid, data.Email, data.Role)
	token, err := jwt.GenerateToken()
	if err != nil {
		responder.Unauthorized("Failed to generate token", err.Error())
		return
	}

	result := models.User{
		Uuid:      data.Uuid,
		Username:  data.Username,
		Email:     data.Email,
		CreatedAt: data.CreatedAt,
	}

	responder.LoginSuccess("Login success", result, token)
}
