package handler

import (
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/services/shipping/internal/usecase"
	"github.com/labstack/echo/v4"
)

type ShippingHandler struct {
	uc *usecase.ShippingUsecase
}

func NewShippingHandler(uc *usecase.ShippingUsecase) *ShippingHandler {
	return &ShippingHandler{uc: uc}
}

// CreateShipping swaggo annotation.
//
//	@Summary		Create shipping
//	@Description	Create new customer order shipping
//	@Tags			shipping-service
//	@Accept			json
//	@Produce		json
//	@Param			CreateShippingRequest	body		entity.CreateShippingRequest	true	"Create shipping request"
//	@Success		200						{object}	entity.Response
//	@Failure		400						{object}	entity.Response
//	@Failure		401						{object}	entity.Response
//	@Failure		404						{object}	entity.Response
//	@Failure		500						{object}	entity.Response
//	@Router			/shipping [post]
func (h *ShippingHandler) CreateShipping(ctx echo.Context) error {
	var req entity.CreateShippingRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	shipping, err := h.uc.CreateShipping(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(shipping))
}

// UpdateShippingStatus swaggo annotation. annotation.
//
//	@Summary		Update shipping
//	@Description	Update customer order shipping status by ID
//	@Tags			shipping-service
//	@Accept			json
//	@Produce		json
//	@Param			shippingId					path		integer								true	"Shipping ID"
//	@Param			UpdateShippingStatusRequest	body		entity.UpdateShippingStatusRequest	true	"Update shipping request"
//	@Success		200							{object}	entity.Response
//	@Failure		400							{object}	entity.Response
//	@Failure		401							{object}	entity.Response
//	@Failure		404							{object}	entity.Response
//	@Failure		500							{object}	entity.Response
//	@Router			/shipping/{shippingId}/status [put]
func (h *ShippingHandler) UpdateShippingStatus(ctx echo.Context) error {
	var req entity.UpdateShippingStatusRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	err := h.uc.UpdateShippingStatus(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccessNilData())
}
