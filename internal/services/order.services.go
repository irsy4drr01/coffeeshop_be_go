package services

import (
	"context"
	"fmt"

	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/repositories"
)

type OrderServiceInterface interface {
	AddOrderService(ctx context.Context, req models.CreateOrderRequest) (models.CreateOrderResponse, error)
}

type OrderService struct {
	repo repositories.OrderRepoInterface
}

func NewOrderService(repo repositories.OrderRepoInterface) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) AddOrderService(ctx context.Context, req models.CreateOrderRequest) (models.CreateOrderResponse, error) {
	for i, item := range req.Items {
		if item.SizeID == 0 {
			req.Items[i].SizeID = 4
		}
	}

	result, err := s.repo.CreateOrder(ctx, req)
	if err != nil {
		return models.CreateOrderResponse{}, err
	}

	resp := models.CreateOrderResponse{
		OrderID:          result.OrderID,
		UserID:           req.UserID,
		Fullname:         req.Fullname,
		Address:          req.Address,
		Phone:            result.Phone,
		DeliveryMethodID: result.DeliveryMethod.ID,
		DeliveryMethod:   result.DeliveryMethod.Name,
		PaymentMethodID:  result.PaymentMethod.ID,
		PaymentMethod:    result.PaymentMethod.Name,
		Items:            result.Items,
		TotalPurchase:    fmt.Sprintf("Rp. %s", result.TotalPurchase.StringFixed(0)),
		DeliveryFee:      fmt.Sprintf("Rp. %d", result.DeliveryMethod.Fee),
		TaxAmount:        fmt.Sprintf("Rp. %s", result.TaxAmount.StringFixed(0)),
		TotalAmount:      fmt.Sprintf("Rp. %s", result.TotalAmount.StringFixed(0)),
	}

	return resp, nil
}
