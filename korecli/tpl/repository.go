package tpl

func RepoProviderTemplate() []byte {
	return []byte(`package repository

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(New{{ .StructName }}Repository)
`)
}

func RepoTemplate() []byte {
	return []byte(`package repository

import (
	"context"

	"{{ .Mod }}/libs/model"
	"{{ .Mod }}/libs/repository"
)

type {{ .ServiceName }}Repo struct {
}

func New{{ .StructName }}Repository() repository.Greeter {
	return &{{ .ServiceName }}Repo{}
}

func (r *{{ .ServiceName }}Repo) Get(ctx context.Context) (*model.Greet, error) {
	return &model.Greet{}, nil
}
`)
}
