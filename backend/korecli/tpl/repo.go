package tpl

func RepoProviderTemplate() []byte {
	return []byte(`package repo

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(New{{ .StructName }}Repository)
`)
}

func RepoTemplate() []byte {
	return []byte(`package repo

import (
	"context"

	"{{ .GoMod }}/pkg/model"
	"{{ .GoMod }}/pkg/repo"
)

type {{ .ServiceName }}Repo struct {
}

func New{{ .StructName }}Repository() repo.GreeterRepository {
	return &{{ .ServiceName }}Repo{}
}

func (r *{{ .ServiceName }}Repo) Get(ctx context.Context) (*model.Greet, error) {
	return &model.Greet{}, nil
}
`)
}
