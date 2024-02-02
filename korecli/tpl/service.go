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
	"context"

	"{{ .Mod }}/libs/proto/v1"
	"{{ .Mod }}/services/{{ .ServiceName }}/usecase"
)

type {{ .StructName }}Service struct {
	v1.Unimplemented{{ .StructName }}Server

	uc *usecase.GreetUsecase
}

func New{{ .StructName }}Service(uc *usecase.GreetUsecase) *{{ .StructName }}Service {
	return &{{ .StructName }}Service{uc: uc}
}

func (s *{{ .StructName }}Service) Greet(ctx context.Context, in *v1.{{ .StructName }}Request) (*v1.{{ .StructName }}Reply, error) {
	greet, err := s.uc.Greet(ctx, in)
	if err != nil {
		return nil, err
	}
	return greet, nil
}
`)
}
