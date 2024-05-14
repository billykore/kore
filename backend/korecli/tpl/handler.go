package tpl

func ServiceProviderTemplate() []byte {
	return []byte(`package service

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(New{{ .StructName }}Handler)
`)
}

func HandlerTemplate() []byte {
	return []byte(`package handler

import (
	"{{ .GoMod }}/pkg/entity"
	"{{ .GoMod }}/services/{{ .ServiceName }}/usecase"
	"github.com/labstack/echo/v4"
)

type {{ .StructName }}Handler struct {
	uc *usecase.{{ .StructName }}Usecase
}

func New{{ .StructName }}Handler(uc *usecase.{{ .StructName }}Usecase) *{{ .StructName }}Service {
	return &{{ .StructName }}Handler{uc: uc}
}

func (s *{{ .StructName }}Handler) Greet(ctx echo.Context) error {
	in := new(entity.{{ .StructName }}Request)
	err := ctx.Bind(in)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	greet, err := s.uc.Greet(ctx.Request().Context(), in)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess("greet", greet))
}`)
}
