package service

import (
	"github.com/billykore/kore/libs/entity"
	"github.com/billykore/kore/services/book/usecase"
	"github.com/labstack/echo/v4"
)

type BookService struct {
	uc *usecase.BookUsecase
}

func NewBookService(uc *usecase.BookUsecase) *BookService {
	return &BookService{uc: uc}
}

func (s *BookService) Greet(ctx echo.Context) error {
	in := new(entity.BookRequest)
	err := ctx.Bind(in)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	greet, err := s.uc.Greet(ctx.Request().Context(), in)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(greet))
}