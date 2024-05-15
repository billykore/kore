package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/repo"
)

type shippingRepo struct {
}

func NewShippingRepository() repo.GreeterRepository {
	return &shippingRepo{}
}

func (r *shippingRepo) Get(ctx context.Context) (*model.Greet, error) {
	return &model.Greet{}, nil
}
