package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/services"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
)

type OrderHandlers struct {
	service services.OrderServiceInterface
}

func NewOrder(service services.OrderServiceInterface) *OrderHandlers {
	return &OrderHandlers{service: service}
}

func (h *OrderHandlers) AddOrderHandler(ctx *gin.Context) {
	res := pkg.NewResponse(ctx)

	var req models.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res.BadRequest("Invalid request", err.Error())
		return
	}

	req.UserID = ctx.GetString(pkg.ContextUserID)
	if req.UserID == "" {
		res.Unauthorized("Unauthorized", "missing user_id in context")
		return
	}

	resp, err := h.service.AddOrderService(ctx, req)
	if err != nil {
		res.InternalServerError("Failed to create order", err.Error())
		return
	}

	res.Created("Order created successfully", resp)
}
