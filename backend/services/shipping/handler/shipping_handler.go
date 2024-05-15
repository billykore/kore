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

func (s *ShippingHandler) Greet(ctx echo.Context) error {
	in := new(entity.ShippingRequest)
	err := ctx.Bind(in)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	greet, err := s.uc.Greet(ctx.Request().Context(), in)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("greet", greet))
}
