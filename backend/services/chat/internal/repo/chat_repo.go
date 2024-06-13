package repo

import (
	"context"

	"github.com/billykore/kore/backend/pkg/model"
)

type ChatRepo struct {
}

func NewChatRepository() *ChatRepo {
	return &ChatRepo{}
}

func (r *ChatRepo) Get(ctx context.Context) (*model.Greet, error) {
	return &model.Greet{}, nil
}
