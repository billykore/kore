package tpl

func RepoProviderTemplate() []byte {
	return []byte(`package repository

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewGreetRepository)
`)
}

func RepoTemplate() []byte {
	return []byte(`package repository

import (
	"context"

	"{{ .Mod }}/libs/model"
	"{{ .Mod }}/libs/repository"
)

type greetRepo struct {
}

func NewGreetRepository() repository.Greeter {
	return &greetRepo{}
}

func (r *greetRepo) Get(ctx context.Context) (*model.Greet, error) {
	return &model.Greet{}, nil
}
`)
}
