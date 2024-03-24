package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
	"github.com/billykore/kore/backend/pkg/repo"
)

type chatRepo struct {
}

func NewChatRepository() repo.GreeterRepository {
	return &chatRepo{}
}

func (r *chatRepo) Get(ctx context.Context) (*model.Greet, error) {
	return &model.Greet{}, nil
}
