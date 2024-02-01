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

	"{{ .Mod }}/services/{{ .ServiceName }}/usecase"
	"{{ .Mod }}/libs/proto/v1"
)

type {{ .StructName }}Service struct {
	v1.UnimplementedGreeterServer

	uc *usecase.GreetUsecase
}

func New{{ .StructName }}Service(uc *usecase.GreetUsecase) *{{ .StructName }}Service {
	return &{{ .StructName }}Service{uc: uc}
}

func (s *{{ .StructName }}Service) Greet(ctx context.Context, in *v1.GreetRequest) (*v1.GreetReply, error) {
	greet, err := s.uc.Greet(ctx, in)
	if err != nil {
		return nil, err
	}
	return greet, nil
}
`)
}
