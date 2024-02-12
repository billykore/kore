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

	"{{ .Mod }}/libs/pkg/log"
	"{{ .Mod }}/libs/proto/v1"
	"{{ .Mod }}/libs/repo"
)

type {{ .StructName }}Usecase struct {
	log  *log.Logger
	repo repo.Greeter
}

func New{{ .StructName }}Usecase(log *log.Logger, repo repository.Greeter) *{{ .StructName }}Usecase {
	return &{{ .StructName }}Usecase{
		log:  log,
		repo: repo,
	}
}

func (uc *{{ .StructName }}Usecase) Greet(ctx context.Context, req *v1.{{ .StructName }}Request) (*v1.{{ .StructName }}Reply, error) {
	uc.log.Usecase("Greet").Infof("Greet %s", req.GetName())
	return &v1.{{ .StructName }}Reply{
		Message: "Hello " + req.GetName(),
	}, nil
}
`)
}
