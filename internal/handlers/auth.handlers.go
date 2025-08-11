package handlers

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/services"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
)

type AuthHandlers struct {
	service services.AuthServiceInterface
}

func NewAuth(service services.AuthServiceInterface) *AuthHandlers {
	return &AuthHandlers{service: service}
}

func (h *AuthHandlers) RegisterHandler(ctx *gin.Context) {
	responder := pkg.NewResponse(ctx)

	// Bind input
	var req models.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responder.BadRequest("Invalid request. Please check your input.", nil)
		return
	}

	res, err := h.service.RegisterService(ctx, req)
	if err != nil {
		errMsg := err.Error()
		log.Printf("[RegisterHandler] Register error: %v", errMsg)

		switch {
		case strings.Contains(errMsg, "already in use"):
			responder.BadRequest("The email is already registered. Please use a different one.", nil)
		case strings.Contains(errMsg, "email is required"):
			responder.BadRequest("Email field cannot be empty.", nil)
		case strings.Contains(errMsg, "invalid email format"):
			responder.BadRequest("The email format is invalid. Please check again.", nil)
		case strings.Contains(errMsg, "password is required"):
			responder.BadRequest("Password field cannot be empty.", nil)
		case strings.Contains(errMsg, "password must"):
			responder.BadRequest(errMsg, nil)
		case strings.Contains(errMsg, "fullname is required"):
			responder.BadRequest("Full name field cannot be empty.", nil)
		case strings.Contains(errMsg, "validation"):
			responder.BadRequest("Some fields are missing or invalid.", nil)
		default:
			responder.InternalServerError("Something went wrong while registering. Please try again later.", nil)
		}
		return
	}

	responder.Created("User registered successfully.", res)
}

func (h *AuthHandlers) LoginHandler(ctx *gin.Context) {
	responder := pkg.NewResponse(ctx)

	var reqBody models.LoginRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		responder.BadRequest("Invalid input!", err.Error())
		return
	}

	responseData, token, err := h.service.LoginService(ctx, reqBody)
	if err != nil {
		errMsg := strings.ToLower(err.Error())

		switch {
		case strings.Contains(errMsg, "invalid input"):
			responder.BadRequest(err.Error(), nil)
		case strings.Contains(errMsg, "wrong email or password"):
			responder.BadRequest("Wrong email or password", nil)
		default:
			log.Printf("[LoginHandler] unexpected error: %v", err)
			responder.InternalServerError("Login failed", nil)
		}
		return
	}

	responder.LoginSuccess("Login successful!", responseData, token)
}
