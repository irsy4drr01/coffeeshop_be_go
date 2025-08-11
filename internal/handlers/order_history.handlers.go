package handlers

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/services"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
)

type OrderHistoryHandlers struct {
	service services.OrderHistoryInterface
}

func NewOrderHistory(service services.OrderHistoryInterface) *OrderHistoryHandlers {
	return &OrderHistoryHandlers{service: service}
}

func (h *OrderHistoryHandlers) FetchAllOrderHistoriesHandler(ctx *gin.Context) {
	responder := pkg.NewResponse(ctx)

	// Ambil userID dari context jwt middleware
	userID := ctx.GetString(pkg.ContextUserID)
	if userID == "" {
		responder.BadRequest("Missing user ID", nil)
		return
	}

	baseHost := ctx.Request.Host
	scheme := "http://"
	if ctx.Request.TLS != nil {
		scheme = "https://"
	}

	orderHistories, err := h.service.FetchAllOrderHistoriesService(ctx.Request.Context(), userID, baseHost, scheme)
	if err != nil {
		responder.InternalServerError("Failed to fetch order histories", err.Error())
		return
	}

	if len(orderHistories) == 0 {
		responder.Success("No order histories found", orderHistories)
		return
	}

	responder.Success("Order histories fetched successfully", orderHistories)
}

func (h *OrderHistoryHandlers) FetchOrderHistoryDetailsHandler(ctx *gin.Context) {
	responder := pkg.NewResponse(ctx)

	userID := ctx.GetString(pkg.ContextUserID)
	if userID == "" {
		responder.BadRequest("Missing user ID", nil)
		return
	}

	orderID := ctx.Param("order_id")
	if orderID == "" {
		responder.BadRequest("Missing order ID", nil)
		return
	}

	baseHost := ctx.Request.Host
	scheme := "http://"
	if ctx.Request.TLS != nil {
		scheme = "https://"
	}

	result, err := h.service.FetchOrderHistoryDetails(ctx.Request.Context(), orderID, userID, baseHost, scheme)
	if err != nil {
		if strings.Contains(err.Error(), "order not found") {
			responder.NotFound("Order not found", nil)
			return
		}
		responder.InternalServerError("Failed to fetch order detail", err.Error())
		return
	}

	responder.Success("Order detail fetched successfully", result)
}
