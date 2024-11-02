package handler

import (
	"github.com/billykore/kore/backend/internal/domain/order"
	"github.com/billykore/kore/backend/internal/infra/messaging/rabbitmq"
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	svc    *order.Service
	rabbit *rabbitmq.Connection
}

func NewOrderHandler(svc *order.Service, rabbit *rabbitmq.Connection) *OrderHandler {
	return &OrderHandler{
		svc:    svc,
		rabbit: rabbit,
	}
}

// Checkout swaggo annotation.
//
//	@Summary		Checkout items
//	@Description	Checkout customer cart items
//	@Tags			order
//	@Accept			json
//	@Produce		json
//	@Param			CheckoutRequest	body		order.CheckoutRequest	true	"Checkout request"
//	@Success		200				{object}	entity.Response
//	@Failure		400				{object}	entity.Response
//	@Failure		401				{object}	entity.Response
//	@Failure		404				{object}	entity.Response
//	@Failure		500				{object}	entity.Response
//	@Router			/orders/checkout [post]
func (h *OrderHandler) Checkout(ctx echo.Context) error {
	var req order.CheckoutRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	err := h.svc.Checkout(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccessNilData())
}

// GetOrderById swaggo annotation.
//
//	@Summary		Get specific order
//	@Description	Get order by ID
//	@Tags			order
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
	var req order.GetRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	res, err := h.svc.GetOrderById(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(res))
}

// PayOrder swaggo annotation.
//
//	@Summary		Pay order
//	@Description	Pay customer order
//	@Tags			order-service
//	@Accept			json
//	@Produce		json
//	@Param			orderId				path		integer					true	"Order ID"
//	@Param			OrderPaymentRequest	body		order.PaymentRequest	true	"Pay order request"
//	@Success		200					{object}	entity.Response
//	@Failure		400					{object}	entity.Response
//	@Failure		401					{object}	entity.Response
//	@Failure		404					{object}	entity.Response
//	@Failure		500					{object}	entity.Response
//	@Router			/orders/{orderId}/pay [post]
func (h *OrderHandler) PayOrder(ctx echo.Context) error {
	var req order.PaymentRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	payment, err := h.svc.PayOrder(ctx.Request().Context(), req)
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
//	@Param			ShippingRequest	body		order.ShippingRequest	true	"Ship order request"
//	@Success		200				{object}	entity.Response
//	@Failure		400				{object}	entity.Response
//	@Failure		401				{object}	entity.Response
//	@Failure		404				{object}	entity.Response
//	@Failure		500				{object}	entity.Response
//	@Router			/orders/{orderId}/ship [post]
func (h *OrderHandler) ShipOrder(ctx echo.Context) error {
	var req order.ShippingRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	shipping, err := h.svc.ShipOrder(ctx.Request().Context(), req)
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
	var req order.CancelOrderRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	err := h.svc.CancelOrder(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccessNilData())
}
