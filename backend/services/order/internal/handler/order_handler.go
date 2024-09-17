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

// Checkout swaggo annotation.
//
//	@Summary		Checkout items
//	@Description	Checkout customer cart items
//	@Tags			order-service
//	@Accept			json
//	@Produce		json
//	@Param			CheckoutRequest	body		entity.CheckoutRequest	true	"Checkout request"
//	@Success		200				{object}	entity.Response
//	@Failure		400				{object}	entity.Response
//	@Failure		401				{object}	entity.Response
//	@Failure		404				{object}	entity.Response
//	@Failure		500				{object}	entity.Response
//	@Router			/orders/checkout [post]
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

// GetOrderById swaggo annotation.
//
//	@Summary		Get specific order
//	@Description	Get order by ID
//	@Tags			order-service
//	@Accept			json
//	@Produce		json
//	@Param			orderId	path		integer	true	"Order ID"
//	@Success		200		{object}	entity.Response
//	@Failure		400		{object}	entity.Response
//	@Failure		401		{object}	entity.Response
//	@Failure		404		{object}	entity.Response
//	@Failure		500		{object}	entity.Response
//	@Router			/orders/{orderId} [get]
func (h *OrderHandler) GetOrderById(ctx echo.Context) error {
	var req entity.OrderRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	order, err := h.uc.GetOrderById(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(order))
}

// PayOrder swaggo annotation.
//
//	@Summary		Pay order
//	@Description	Pay customer order
//	@Tags			order-service
//	@Accept			json
//	@Produce		json
//	@Param			orderId				path		integer						true	"Order ID"
//	@Param			OrderPaymentRequest	body		entity.OrderPaymentRequest	true	"Pay order request"
//	@Success		200					{object}	entity.Response
//	@Failure		400					{object}	entity.Response
//	@Failure		401					{object}	entity.Response
//	@Failure		404					{object}	entity.Response
//	@Failure		500					{object}	entity.Response
//	@Router			/orders/{orderId}/pay [post]
func (h *OrderHandler) PayOrder(ctx echo.Context) error {
	var req entity.OrderPaymentRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	payment, err := h.uc.PayOrder(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(payment))
}

// ShipOrder swaggo annotation.
//
//	@Summary		Ship order
//	@Description	Ship customer order
//	@Tags			order-service
//	@Accept			json
//	@Produce		json
//	@Param			orderId			path		integer					true	"Order ID"
//	@Param			ShippingRequest	body		entity.ShippingRequest	true	"Ship order request"
//	@Success		200				{object}	entity.Response
//	@Failure		400				{object}	entity.Response
//	@Failure		401				{object}	entity.Response
//	@Failure		404				{object}	entity.Response
//	@Failure		500				{object}	entity.Response
//	@Router			/orders/{orderId}/ship [post]
func (h *OrderHandler) ShipOrder(ctx echo.Context) error {
	var req entity.ShippingRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	shipping, err := h.uc.ShipOrder(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(shipping))
}

// CancelOrder swaggo annotation.
//
//	@Summary		Cancel specific order
//	@Description	Cancel customer order by ID
//	@Tags			order-service
//	@Accept			json
//	@Produce		json
//	@Param			orderId	path		integer	true	"Order ID"
//	@Success		200		{object}	entity.Response
//	@Failure		400		{object}	entity.Response
//	@Failure		401		{object}	entity.Response
//	@Failure		404		{object}	entity.Response
//	@Failure		500		{object}	entity.Response
//	@Router			/orders/{orderId} [delete]
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
