package pkg

import (
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type CloudinaryMock struct {
	mock.Mock
}

func (m *CloudinaryMock) UploadFile(ctx *gin.Context, file interface{}, fileName string) (*uploader.UploadResult, error) {
	args := m.Called(ctx, file, fileName)
	return args.Get(0).(*uploader.UploadResult), args.Error(1)
}

func (m *CloudinaryMock) DeleteFile(ctx *gin.Context, publicID string) (*uploader.DestroyResult, error) {
	args := m.Called(ctx, publicID)
	return args.Get(0).(*uploader.DestroyResult), args.Error(1)
}
