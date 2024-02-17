package repo

import (
	"context"

	"github.com/billykore/kore/libs/model"
	"github.com/billykore/kore/libs/repo"
)

type bookRepo struct {
}

func NewBookRepository() repo.GreeterRepository {
	return &bookRepo{}
}

func (r *bookRepo) Get(ctx context.Context) (*model.Greet, error) {
	return &model.Greet{}, nil
}
