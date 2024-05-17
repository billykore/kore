package handler

import (
	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/billykore/kore/backend/services/shipping/usecase"
	"github.com/labstack/echo/v4"
)

type ShippingHandler struct {
	uc *usecase.ShippingUsecase
}

func NewShippingHandler(uc *usecase.ShippingUsecase) *ShippingHandler {
	return &ShippingHandler{uc: uc}
}

func (h *ShippingHandler) CreateShipping(ctx echo.Context) error {
	var req entity.CreateShippingRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(entity.ResponseBadRequest(err))
	}
	resp, err := h.uc.CreateShipping(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("shipping", resp))
}

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
