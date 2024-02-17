package tpl

func ServiceProviderTemplate() []byte {
	return []byte(`package service

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(New{{ .StructName }}Service)
`)
}

func ServiceTemplate() []byte {
	return []byte(`package service

import (
	"{{ .Mod }}/libs/entity"
	"{{ .Mod }}/services/{{ .ServiceName }}/usecase"
	"github.com/labstack/echo/v4"
)

type {{ .StructName }}Service struct {
	uc *usecase.{{ .StructName }}Usecase
}

func New{{ .StructName }}Service(uc *usecase.{{ .StructName }}Usecase) *{{ .StructName }}Service {
	return &{{ .StructName }}Service{uc: uc}
}

func (s *{{ .StructName }}Service) Greet(ctx echo.Context) error {
	in := new(entity.{{ .StructName }}Request)
	err := ctx.Bind(in)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	greet, err := s.uc.Greet(ctx.Request().Context(), in)
	if err != nil {
		return ctx.JSON(entity.ResponseError(err))
	}
	return ctx.JSON(entity.ResponseSuccess(greet))
}`)
}
