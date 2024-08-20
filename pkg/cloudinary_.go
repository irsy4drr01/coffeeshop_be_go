package pkg

import (
	"crypto/rand"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

type CloudinaryInterface interface {
	UploadFile(ctx *gin.Context, file interface{}, fileName string) (*uploader.UploadResult, error)
	DeleteFile(ctx *gin.Context, publicID string) (*uploader.DestroyResult, error)
}

type Cloudinary struct {
	CLD *cloudinary.Cloudinary
}

func NewCloudinaryUtil() *Cloudinary {
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)
	if err != nil {
		log.Fatal("Failed to initiate Cloudinary: %w", err)
	}

	return &Cloudinary{
		CLD: cld,
	}
}

func (c *Cloudinary) UploadFile(ctx *gin.Context, file interface{}, fileName string) (*uploader.UploadResult, error) {
	uploadResult, err := c.CLD.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: fileName,
	})
	if err != nil {
		return nil, err
	}
	return uploadResult, nil
}

func (c *Cloudinary) DeleteFile(ctx *gin.Context, publicID string) (*uploader.DestroyResult, error) {
	destroyResult, err := c.CLD.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})
	if err != nil {
		return nil, err
	}
	return destroyResult, nil
}

func RandomInt(max int) (int, error) {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(nBig.Int64()), nil
}

func GetPublicIDFromURL(url string) string {
	parts := strings.Split(url, "/")
	fileName := parts[len(parts)-1]
	publicID := strings.Split(fileName, ".")[0]
	return publicID
}
