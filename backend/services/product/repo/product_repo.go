package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/repo"
)

type productRepo struct {
}

func NewProductRepository() repo.GreeterRepository {
	return &productRepo{}
}

func (r *productRepo) Get(ctx context.Context) (*model.Greet, error) {
	return &model.Greet{}, nil
}
