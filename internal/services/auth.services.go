package services

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
	"github.com/irsy4drr01/coffeeshop_be_go/utils"
)

type AuthServiceInterface interface {
	RegisterService(ctx context.Context, req models.RegisterRequest) (*models.RegisterResponse, error)
	LoginService(ctx context.Context, req models.LoginRequest) (*models.LoginResponseData, string, error)
}

type AuthService struct {
	repo repositories.AuthRepoInterface
}

func NewAuthService(repo repositories.AuthRepoInterface) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) RegisterService(ctx context.Context, req models.RegisterRequest) (*models.RegisterResponse, error) {
	// Trim
	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)
	req.Fullname = strings.TrimSpace(req.Fullname)

	// Govalidator
	if _, err := govalidator.ValidateStruct(req); err != nil {
		return nil, fmt.Errorf(utils.ConvertValidatorError(err.Error()))
	}

	// Email format
	if err := utils.ValidateEmailFormat(req.Email); err != nil {
		return nil, err
	}

	// Password strength
	if err := utils.ValidatePasswordStrength(req.Password); err != nil {
		return nil, err
	}

	// Hash
	hashedPwd, err := pkg.HashPassword(req.Password)
	if err != nil {
		log.Printf("[AuthService][Register] HashPassword error: %v", err)
		return nil, fmt.Errorf("failed to process password")
	}

	// Prepare models
	user := &models.UserAuth{
		Email:    req.Email,
		Password: hashedPwd,
	}
	profile := &models.ProfileAuth{
		Fullname: req.Fullname,
	}

	// Call repo
	newUser, newProfile, err := s.repo.CreateUserAndProfile(ctx, user, profile)
	if err != nil {
		if strings.Contains(err.Error(), "email already in use") {
			return nil, fmt.Errorf("email already in use")
		}
		log.Printf("[AuthService][Register] CreateUserAndProfile error: %v", err)
		return nil, fmt.Errorf("failed to register user")
	}

	// Response
	res := &models.RegisterResponse{
		ID:       newUser.ID,
		Email:    newUser.Email,
		Fullname: newProfile.Fullname,
		Created:  newUser.CreatedAt,
	}

	return res, nil
}

func (s *AuthService) LoginService(ctx context.Context, req models.LoginRequest) (*models.LoginResponseData, string, error) {
	// Govalidator
	if _, err := govalidator.ValidateStruct(&req); err != nil {
		return nil, "", fmt.Errorf("invalid input: %s", utils.ConvertValidatorError(err.Error()))
	}

	// Email format
	if err := utils.ValidateEmailFormat(req.Email); err != nil {
		return nil, "", fmt.Errorf("invalid input: %s", err.Error())
	}

	// Get user
	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		log.Printf("[LoginService][GetByEmail] error: %v", err)
		return nil, "", fmt.Errorf("login failed")
	}
	if user == nil {
		return nil, "", fmt.Errorf("wrong email or password")
	}

	// Verify password
	if err := pkg.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, "", fmt.Errorf("wrong email or password")
	}

	// JWT
	jwt := pkg.NewJWT(user.ID, user.Email, user.Role, user.IsVerified)
	token, err := jwt.GenerateToken()
	if err != nil {
		log.Printf("[AuthService][Login] GenerateToken error: %v", err)
		return nil, "", fmt.Errorf("login failed")
	}

	// Build response data
	res := &models.LoginResponseData{
		ID:      user.ID,
		Email:   user.Email,
		Role:    user.Role,
		Created: user.CreatedAt,
	}

	return res, token, nil
}
