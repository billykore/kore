package handler

import (
	"github.com/billykore/kore/backend/internal/domain/shipping"
	"github.com/billykore/kore/backend/pkg/entity"

	"github.com/labstack/echo/v4"
)

type ShippingHandler struct {
	svc *shipping.Service
}

func NewShippingHandler(svc *shipping.Service) *ShippingHandler {
	return &ShippingHandler{svc: svc}
}

// CreateShipping swaggo annotation.
//
//	@Summary		Create shipping
//	@Description	Create new customer order shipping
//	@Tags			shipping
//	@Accept			json
//	@Produce		json
//	@Param			CreateShippingRequest	body		shipping.CreateShippingRequest	true	"Create shipping request"
//	@Success		200						{object}	entity.Response
//	@Failure		400						{object}	entity.Response
//	@Failure		401						{object}	entity.Response
//	@Failure		404						{object}	entity.Response
//	@Failure		500						{object}	entity.Response
//	@Router			/shipping [post]
func (h *ShippingHandler) CreateShipping(ctx echo.Context) error {
	var req shipping.CreateShippingRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	res, err := h.svc.CreateShipping(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(res))
}

// UpdateShippingStatus swaggo annotation. annotation.
//
//	@Summary		Update shipping
//	@Description	Update customer order shipping status by ID
//	@Tags			shipping
//	@Accept			json
//	@Produce		json
//	@Param			shippingId					path		integer									true	"Shipping ID"
//	@Param			UpdateShippingStatusRequest	body		shipping.UpdateShippingStatusRequest	true	"Update shipping request"
//	@Success		200							{object}	entity.Response
//	@Failure		400							{object}	entity.Response
//	@Failure		401							{object}	entity.Response
//	@Failure		404							{object}	entity.Response
//	@Failure		500							{object}	entity.Response
//	@Router			/shipping/{shippingId}/status [put]
func (h *ShippingHandler) UpdateShippingStatus(ctx echo.Context) error {
	var req shipping.UpdateShippingStatusRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	err := h.svc.UpdateShippingStatus(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccessNilData())
}
