package services

import (
	"context"
	"fmt"

	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
	"github.com/irsy4drr01/coffeeshop_be_go/utils"
)

type ProductServiceInterface interface {
	FetchAllProductsService(ctx context.Context, queryParams models.ProductQueryParams, baseHost, scheme, baseURL string) (models.ProductsResponse, pkg.Pagination, error)
	FetchProductDetailsService(ctx context.Context, productID, baseHost, scheme string) (models.ProductDetailsResponse, error)
}

type ProductService struct {
	repo repositories.ProductRepoInterface
}

func NewProductService(repo repositories.ProductRepoInterface) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) FetchAllProductsService(ctx context.Context, queryParams models.ProductQueryParams, baseHost, scheme, baseURL string) (models.ProductsResponse, pkg.Pagination, error) {
	// default pagination
	if queryParams.Page < 1 {
		queryParams.Page = 1
	}
	if queryParams.Limit < 1 {
		queryParams.Limit = 8
	}

	products, total, err := s.repo.GetAllProducts(ctx, queryParams)
	if err != nil {
		return nil, pkg.Pagination{}, err
	}

	// Format response
	var response models.ProductsResponse
	for _, p := range products {
		imgURL := utils.BuildImageURL(p.ProductImg, baseHost, scheme)
		priceStr, finalPriceStr := utils.CalculatePriceAndFinal(p.Price, p.IsDiscount, p.DiscountRate)

		response = append(response, models.ProductResponse{
			ID:           p.ProductID,
			ProductName:  p.ProductName,
			CategoryID:   p.CategoryID,
			CategoryName: p.CategoryName,
			ProductImg:   imgURL,
			Description:  p.Description,
			TotalSold:    p.TotalSold,
			TotalLike:    p.TotalLikes,
			Price:        priceStr,
			FinalPrice:   finalPriceStr,
			IsDiscount:   p.IsDiscount,
			IsDeleted:    p.IsDeleted,
		})
	}

	// Build pagination meta
	totalPage := (total + queryParams.Limit - 1) / queryParams.Limit

	meta := pkg.Pagination{
		TotalData: total,
		TotalPage: totalPage,
		Page:      queryParams.Page,
	}

	// Build prev and next links
	baseQuery := utils.BuildQueryString(queryParams)

	if queryParams.Page > 1 {
		prevPage := queryParams.Page - 1
		meta.PrevLink = fmt.Sprintf("%s?%s&page=%d&limit=%d", baseURL, baseQuery, prevPage, queryParams.Limit)
	}
	if queryParams.Page < totalPage {
		nextPage := queryParams.Page + 1
		meta.NextLink = fmt.Sprintf("%s?%s&page=%d&limit=%d", baseURL, baseQuery, nextPage, queryParams.Limit)
	}

	return response, meta, nil
}

func (s *ProductService) FetchProductDetailsService(ctx context.Context, productID, baseHost, scheme string) (models.ProductDetailsResponse, error) {
	product, slots, sizes, err := s.repo.GetOneProduct(ctx, productID)
	if err != nil {
		return models.ProductDetailsResponse{}, fmt.Errorf("not found")
	}

	// --- Build slot_number logic: default + override
	defaultImg := "product_default.webp"
	slotMap := map[int]string{
		1: defaultImg,
		2: defaultImg,
		3: defaultImg,
	}
	for _, slot := range slots {
		slotMap[slot.SlotNumber] = slot.ImageFile
	}

	imgResp := models.ProductImgResponse{}
	for i := 1; i <= 3; i++ {
		img := slotMap[i]
		imgURL := utils.BuildImageURL(img, baseHost, scheme)
		switch i {
		case 1:
			imgResp.Img1 = imgURL
		case 2:
			imgResp.Img2 = imgURL
		case 3:
			imgResp.Img3 = imgURL
		}
	}

	priceStr, finalPriceStr := utils.CalculatePriceAndFinal(product.Price, product.IsDiscount, product.DiscountRate)

	response := models.ProductDetailsResponse{
		ProductID:   product.ProductID,
		ProductName: product.ProductName,
		ProductImgs: imgResp,
		Price:       priceStr,
		FinalPrice:  finalPriceStr,
		TotalSold:   product.TotalSold,
		TotalLike:   product.TotalLikes,
		Description: product.Description,
		IsDiscount:  product.IsDiscount,
	}

	if product.CategoryID == 1 || product.CategoryID == 2 {
		response.DataSize = sizes
	} else {
		response.DataSize = []models.Size{}
	}

	return response, nil
}
