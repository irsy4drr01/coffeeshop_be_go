package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
	"github.com/irsy4drr01/coffeeshop_be_go/utils"
)

type OrderHistoryInterface interface {
	FetchAllOrderHistoriesService(ctx context.Context, userID, baseHost, scheme string) (models.OrderHistoriesRes, error)
	FetchOrderHistoryDetails(ctx context.Context, orderID, userID, baseHost, scheme string) (models.OrderHistoryDetailRes, error)
}

type OrderHistoryService struct {
	repo repositories.OrderHistoryRepoInterface
}

func NewOrderHistoryService(repo repositories.OrderHistoryRepoInterface) *OrderHistoryService {
	return &OrderHistoryService{repo: repo}
}

func (s *OrderHistoryService) FetchAllOrderHistoriesService(ctx context.Context, userID, baseHost, scheme string) (models.OrderHistoriesRes, error) {
	orderHistoriesDB, err := s.repo.GetAllOrderHistories(ctx, userID)
	if err != nil {
		return nil, err
	}

	var orderHistoriesRes models.OrderHistoriesRes
	for _, o := range orderHistoriesDB {
		orderHistoriesRes = append(orderHistoriesRes, models.OrderHistoryRes{
			ID:          o.ID,
			Image:       utils.BuildImageURL(o.Image, baseHost, scheme),
			Status:      o.Status,
			CreatedAt:   utils.FormatDate(o.CreatedAt),
			TotalAmount: utils.FormatRupiah(o.TotalAmount),
		})
	}

	return orderHistoriesRes, nil
}

func (s *OrderHistoryService) FetchOrderHistoryDetails(ctx context.Context, orderID, userID, baseHost, scheme string) (models.OrderHistoryDetailRes, error) {
	dbDetail, err := s.repo.GetOrderHistoryDetails(ctx, orderID, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.OrderHistoryDetailRes{}, errors.New("order not found")
		}
		return models.OrderHistoryDetailRes{}, err
	}

	resItems := make(models.OrderHistoryItemsRes, 0, len(dbDetail.Items))
	for _, item := range dbDetail.Items {
		resItems = append(resItems, models.OrderHistoryItemRes{
			ProductName:  item.ProductName,
			Image:        utils.BuildImageURL(item.Image, baseHost, scheme),
			Qty:          item.Qty,
			Size:         item.Size,
			Temperature:  utils.FormatTemperature(item.IsIced, item.CategoryID),
			BasePrice:    utils.FormatRupiah(item.BasePrice),
			FinalPrice:   utils.FormatRupiah(item.FinalPrice),
			DiscountName: item.DiscountName,
		})
	}

	result := models.OrderHistoryDetailRes{
		ID:                dbDetail.ID,
		FullName:          dbDetail.FullName,
		Address:           dbDetail.Address,
		Phone:             dbDetail.Phone,
		PaymentMethod:     dbDetail.PaymentMethod,
		Status:            dbDetail.Status,
		TotalPurchase:     utils.FormatRupiah(dbDetail.TotalPurchase),
		DeliveryMethod:    dbDetail.DeliveryMethod,
		DeliveryMethodFee: utils.FormatRupiah(dbDetail.DeliveryMethodFee),
		Tax:               utils.FormatRupiah(dbDetail.Tax),
		TotalAmount:       utils.FormatRupiah(dbDetail.TotalAmount),
		CreatedAt:         utils.FormatDate(dbDetail.CreatedAt),
		Items:             resItems,
	}

	return result, nil
}
