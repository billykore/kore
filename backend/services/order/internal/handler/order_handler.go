package handler

import (
	"context"

	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/pkg/net/rabbit"
	"github.com/billykore/kore/backend/services/order/internal/usecase"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	uc     *usecase.OrderUsecase
	rabbit *rabbit.Rabbit
}

func NewOrderHandler(uc *usecase.OrderUsecase, rabbit *rabbit.Rabbit) *OrderHandler {
	return &OrderHandler{
		uc:     uc,
		rabbit: rabbit,
	}
}

func (h *OrderHandler) Checkout(ctx echo.Context) error {
	var req entity.CheckoutRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	err := h.uc.Checkout(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccessNilData())
}

func (h *OrderHandler) GetOrderById(ctx echo.Context) error {
	var req entity.OrderRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	order, err := h.uc.GetOrderById(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("order", order))
}

func (h *OrderHandler) PayOrder(ctx echo.Context) error {
	var req entity.OrderPaymentRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	payment, err := h.uc.PayOrder(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("payment", payment))
}

func (h *OrderHandler) ShipOrder(ctx echo.Context) error {
	var req entity.ShippingRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	shipping, err := h.uc.ShipOrder(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("shipping", shipping))
}

func (h *OrderHandler) CancelOrder(ctx echo.Context) error {
	var req entity.CancelOrderRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	err := h.uc.CancelOrder(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccessNilData())
}

func (h *OrderHandler) ListenOrderStatusChanges() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	h.rabbit.Consume(ctx, h.uc.ListenOrderStatusChanges)
}
