package tpl

func UsecaseProviderTemplate() []byte {
	return []byte(`package usecase

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(New{{ .StructName }}Usecase)
`)
}

func UsecaseTemplate() []byte {
	return []byte(`package usecase

import (
	"context"

	"{{ .GoMod }}/pkg/entity"
	"{{ .GoMod }}/pkg/log"
	"{{ .GoMod }}/pkg/repo"
)

type {{ .StructName }}Usecase struct {
	log  *log.Logger
	repo repo.GreeterRepository
}

func New{{ .StructName }}Usecase(log *log.Logger, repo repo.GreeterRepository) *{{ .StructName }}Usecase {
	return &{{ .StructName }}Usecase{
		log:  log,
		repo: repo,
	}
}

func (uc *{{ .StructName }}Usecase) Greet(ctx context.Context, req *entity.{{ .StructName }}Request) (*entity.{{ .StructName }}Response, error) {
	uc.log.Usecase("Greet").Infof("Greet %s", req.Name)
	return &entity.{{ .StructName }}Response{
		Message: "Hello " + req.Name,
	}, nil
}
`)
}
