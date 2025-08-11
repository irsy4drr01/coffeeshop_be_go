package services

import (
	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
	"github.com/irsy4drr01/coffeeshop_be_go/utils"
)

type UserServiceInterface interface {
	FetchProfileService(ctx *gin.Context)
	// PatchProfileService(ctx context.Context, ......, ......) error
}

type UserService struct {
	repo repositories.UserRepoInterface
}

func NewUserService(repo repositories.UserRepoInterface) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) FetchProfileService(ctx *gin.Context) {
	responder := pkg.NewResponse(ctx)
	userID := ctx.GetString(pkg.ContextUserID)
	user, err := s.repo.GetOneUser(ctx, userID)
	if err != nil {
		responder.NotFound("User not found", nil)
		return
	}

	// Build profile image URL
	baseHost := ctx.Request.Host
	scheme := "http://"
	if ctx.Request.TLS != nil {
		scheme = "https://"
	}

	profileImage := utils.BuildImageProfileURL(user.Image, baseHost, scheme)

	formattedDate := utils.FormatDate(user.CreatedAt)
	createdAt := "since " + formattedDate

	res := models.UserProfileResponse{
		FullName:   user.FullName,
		ProfileImg: profileImage,
		Email:      user.Email,
		Phone:      user.Phone,
		Address:    user.Address,
		CreatedAt:  createdAt,
	}

	responder.Success("Profile fetched successfully", res)
}

// func (s *UserService) PatchProfileService(ctx context.Context, ......, ......) error {

// }
